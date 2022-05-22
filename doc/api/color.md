# Colors API Reference

To find and match colors

### `IsColor()`

* `color1` [color.Color]()
* `color2` [color.Color]()
* `similarity` [float64]()

Returns [bool]()

Check if the two colors are similar to the `similarity` threshold.

---

### `IsColorWithOffset()`

* `color1` [color.Color]()
* `color2` [color.Color]()
* `offset` [color.Color]()
* `similarity` [float64]()

Returns [bool]()

Check if the two colors are similar to the `similarity` threshold with additional offset(noise) in color2.

---

### `FindColor()`

* `src` [color.Color]()
* `colorToFind` [color.Color]()
* `similarity` [float64]()

Returns
* `x` [int]()
* `y` [int]()

Find the first coordinates of a color within an image.

---

### `FindColorWithOffset()`

* `src` [color.Color]()
* `colorToFind` [color.Color]()
* `offset` [color.Color]()
* `similarity` [float64]()

Returns
* `x` [int]()
* `y` [int]()

Find the first coordinates of a color with offset(noise) within an image.

---

### `GetColorNum()`

* `img` [color.Color]()
* `colorToFind` [color.Color]()
* `similarity` [float64]()

Returns
* `count` [uint]()
* `error` [error]()

Find the occurrence of a color within an image.

---

### `GetColorNumWithOffset()`

* `img` [color.Color]()
* `colorToFind` [color.Color]()
* `offset` [color.Color]()
* `similarity` [float64]()

Returns
* `count` [uint]()
* `error` [error]()

Find the occurrence of a color with offset(noise) within an image.

---

### `RGBToHex()`

* `src` [color.Color]()

Returns [string]()

Return the RGB value as Hex.

---

### `HexToRGB()`

* `hex` [string]()

Returns [color.RGBA]()

Return the Hex value as RGBA.
