# OCR API Reference

To extract text from an image

### `InitOCR()`

* `language` [OCRLangugae]()

Returns [void]()

Needs to be call to initialize OCR.

---

### `FindTextWithBytes()`

* `imgBytes` [[]byte]()

Returns
* [string]()
* [error]()

Find text within a slice of bytes containing the image.

---

### `FindTextWithFile()`

* `fileName` [string]()
Returns
* [image.Image]()
* [error]()

Find text with a file containing the image.

---

### `CloseOCR()`

Returns [void]()

Needs to be call after being done with OCR.
