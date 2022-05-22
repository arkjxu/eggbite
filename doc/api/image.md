# Image API Reference

To find and match image

### `FindImage()`

* `src` [image.Image]()
* `toFind` [image.Image]()

Returns
* `x` [int]()
* `y` [int]()

Find location of image in src image.

---

### `GetImageSimilarity()`

* `src` [image.Image]()
* `compareWith` [image.Image]()
* `ht` [HashType]()

Returns
* [float64]()
* [error]()

Find the similarity in percentages between two image.

---

### `ImageFromBytes()`

* `imgBytes` [[]byte]()

Returns
* [image.Image]()
* [error]()

Get an image.Image from slice of bytes containing an image.

---

### `ImageFromFile()`

* `filename` [string]()

Returns
* [image.Image]()
* [error]()

Get an image.Image from a file containing an image.
