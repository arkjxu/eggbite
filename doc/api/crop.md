# Crop API Reference

To crop an image

### `CropBytes()`

* `imgBytes` [[]byte]()
* `x1` [int]()
* `y1` [int]()
* `x2` [int]()
* `y2` [int]()

Returns
* [image.Image]()
* [error]()

Crop an image using bytes slices containing the image

---

### `CropFile()`

* `filename` [string]()
* `x1` [int]()
* `y1` [int]()
* `x2` [int]()
* `y2` [int]()

Returns
* [image.Image]()
* [error]()

Crop an image using a file containing the image

---

### `CropImage()`

* `img` [image.Image]()
* `x1` [int]()
* `y1` [int]()
* `x2` [int]()
* `y2` [int]()

Returns
* [image.Image]()
* [error]()

Crop an image using an image

---