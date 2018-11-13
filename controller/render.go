package controller

/*
#cgo CFLAGS: -I../indigo/include
#cgo LDFLAGS: -L../indigo/lib -lindigo -lindigo-renderer
#include <indigo-renderer.h>
#include <indigo.h>
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"../helper"
	"log"
	"net/http"
	"strconv"
	"unsafe"
)

func RenderStructureController(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	smilesString := query.Get("smiles")
	renderType := query.Get("type")
	renderWidth := 100
	renderHeight := 100
	background := "256, 256, 256"

	contentType := "image/svg+xml"

	if smilesString == "" {
		http.Error(w, "Cannot find SMILES parameter", http.StatusBadRequest)
		return
	}

	if rw := query.Get("width"); rw != "" {
		i, err := strconv.Atoi(rw)
		if err != nil {
			log.Fatalln(err)
			return
		} else if i < 100 {
			i = 100
		}
		renderWidth = i
	}

	if rh := query.Get("height"); rh != "" {
		i, err := strconv.Atoi(rh)
		if err != nil {
			log.Fatalln(err)
			return
		} else if i < 100 {
			i = 100
		}
		renderHeight = i
	}

	if bg := query.Get("background"); bg != "" {
		background = bg
	}

	if renderType == "" || renderType == "svg" {
		renderType = "svg"
		contentType = "image/svg+xml"
	} else {
		renderType = "png"
		contentType = "image/png"
	}

	err, pngImage := renderStructure(smilesString, renderType, renderWidth, renderHeight, background)

	if err != nil {
		log.Printf("\nError\n%s", err.Error())
	}

	w.Header().Set("Content-Type", contentType)

	w.Write(pngImage)
}

func renderStructure(smiles string, typeRendering string, width int, height int, background string) (error, []byte) {
	var result C.int
	var length C.int

	indigoWriteBuffer := C.indigoWriteBuffer()

	renderOption := C.CString("render-output-format")
	renderMargin := C.CString("render-margins")
	renderImageWidth := C.CString("render-image-width")
	renderImageHeight := C.CString("render-image-height")
	renderColoring := C.CString("render-coloring")
	renderBackgroundColor := C.CString("render-background-color")

	renderType := C.CString(typeRendering)
	imageBackground := C.CString(background)

	smilesC := C.CString(smiles)
	molecule := C.indigoLoadMoleculeFromString(smilesC)
	C.indigoDearomatize(molecule)

	C.indigoSetOption(renderOption, renderType)
	C.indigoSetOptionXY(renderMargin, 5, 5)
	C.indigoSetOptionInt(renderImageWidth, C.int(width))
	C.indigoSetOptionInt(renderImageHeight, C.int(height))
	C.indigoSetOptionBool(renderColoring, 1)
	C.indigoSetOption(renderBackgroundColor, imageBackground)

	result = C.indigoRender(molecule, indigoWriteBuffer)

	if result != 1 {
		log.Fatal("cannot render molecule")
		return helper.Error{Message: "Cannot render molecule"}, nil
	}

	var stringBuffer *C.char
	var byteBuffer []byte

	C.indigoToBuffer(indigoWriteBuffer, &stringBuffer, &length)

	byteBuffer = C.GoBytes(unsafe.Pointer(stringBuffer), length)

	defer func() {
		C.indigoFree(indigoWriteBuffer)
		C.indigoFree(molecule)
		C.indigoFreeAllObjects()

		C.free(unsafe.Pointer(renderMargin))
		C.free(unsafe.Pointer(renderImageWidth))
		C.free(unsafe.Pointer(renderImageHeight))
		C.free(unsafe.Pointer(renderColoring))
		C.free(unsafe.Pointer(renderBackgroundColor))

		C.free(unsafe.Pointer(smilesC))
		C.free(unsafe.Pointer(imageBackground))
		C.free(unsafe.Pointer(renderType))
		C.free(unsafe.Pointer(renderOption))
	}()
	return nil, byteBuffer
}
