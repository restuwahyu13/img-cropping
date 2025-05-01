## Take-Home Test: Image Cropping via Border Detection

### 🌟 Objective

You are given an image file `image.png` that contains a rectangular bordered area. Your task is to **programmatically crop the image**, extracting only the section that is enclosed by the **black border** (the border itself must still exist in the result).

### 🛠️ Requirements

- **Input**: `image.png` (provided in the same folder as your code [Click here to view the image](./image.png))
- **Output**: A cropped image saved as `output.png` containing **only** the bordered section (the border is still visible).
- **Method**: You must detect the border **by reading the pixel colors** (e.g., detect black pixels).

---

### ❌ Restrictions

- **Do not use any third-party libraries**.
- You are only allowed to use **Go's standard library**, especially the `image`, `image/color`, and `image/png` packages.
- The solution **must use manual pixel inspection** to find the crop boundaries.

---

### ✅ Bonus (Optional)

- Save a `.log` file or `.txt` file that logs each coordinate detected as part of the border area.

---

### 💡 Hints

- Loop through the image pixels using nested `for` loops.
- A pixel is considered black if `R=0`, `G=0`, and `B=0`.
- Track the minimum and maximum `x` and `y` values of black pixels to get the bounding box.
