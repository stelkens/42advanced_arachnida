*This project has been created as part of the 42 advanced cybersecurity piscine by tstelken.*

# 🕷️ Spider & 🦂 Scorpion

A website image crawler and metadata extraction toolkit written in **Go (Golang)**.

This project consists of two independent command-line programs, **Spider** and **Scorpion**, designed to work with images found on websites.

- **Spider** crawls websites and downloads images recursively.
- **Scorpion** analyzes image files and extracts their metadata, including EXIF information.

Together, they provide a simple workflow for collecting images from websites and inspecting their embedded information.

---

## 🕷️ Spider

Spider is a recursive website image crawler that extracts and downloads images from a given URL.

By providing a website URL as a parameter, Spider searches for images and downloads them locally. When recursive mode is enabled, it follows links to other pages of the same domain and continues searching for additional images.

### Features

- Extracts images from websites.
- Recursively crawls linked pages of the same domain.
- Configurable recursion depth.
- Custom download directory.
- Supports multiple image formats.
- Command-line argument handling.

### Supported Image Formats

Spider downloads the following image extensions by default:

- `.jpg`
- `.jpeg`
- `.png`
- `.gif`
- `.bmp`

### Usage

```bash
./spider [-rlp] URL
```

### Available Options

| Option      | Description                                                        |
| ----------- | ------------------------------------------------------------------ |
| `-r`        | Enables recursive image downloading.                               |
| `-l [N]`    | Sets the maximum recursion depth. Default: `5`.                    |
| `-p [PATH]` | Specifies the directory for downloaded images. Default: `./data/`. |

### Examples

Download images from a website:

```bash
./spider https://example.com
```

Recursively download images:

```bash
./spider -r https://example.com
```

Recursively download images with a maximum depth of 3:

```bash
./spider -r -l 3 https://example.com
```

Specify a custom download directory:

```bash
./spider -p ./images/ https://example.com
```

Combine recursive downloading, depth limitation, and a custom directory:

```bash
./spider -r -l 3 -p ./downloads/ https://example.com
```

---

## 🦂 Scorpion

Scorpion is an image metadata analyzer that parses image files and extracts information stored within them.

It accepts one or multiple image files as command-line arguments and displays their metadata directly in the terminal.

The main purpose is to inspect images and retrieve useful information about their origin and properties.

### Features

- Analyzes multiple image files in a single execution.
- Extracts basic image attributes.
- Reads EXIF metadata when available.
- Displays image creation dates when available.
- Processes supported image formats.
- Presents extracted information directly in the terminal.

### Supported Image Formats

- `.jpg`
- `.jpeg`
- `.png`
- `.gif`
- `.bmp`

### Usage

```bash
./scorpion FILE1 [FILE2 ...]
```

### Examples

Analyze a single image:

```bash
./scorpion image.jpg
```

Analyze multiple images:

```bash
./scorpion image1.jpg image2.png image3.gif
```

Analyze images downloaded by Spider:

```bash
./scorpion ../spider/data/*
```

### Metadata Extraction

Depending on the image format and available metadata, Scorpion can display information such as:

- Image creation date.
- EXIF metadata.
- Camera information.
- Image properties.
- Other embedded metadata.

---

## ⚙️ Installation

Make sure you have [Go](https://go.dev/dl/) installed.

Check your installation:

```bash
go version
```

Clone the repository:

```bash
git clone <repository-url> spider-scorpion-42
cd spider-scorpion-42
```

Build the programs:

```bash
cd spider/
go build -o spider .
cd ../scorpion/
go build -o scorpion .
```

---

## 🔄 Workflow

Spider and Scorpion can be used together as a simple image collection and analysis pipeline.

1. **Spider** receives a website URL.
2. It crawls the website and downloads the discovered images.
3. The images are stored in the specified output directory.
4. **Scorpion** receives the downloaded image files.
5. It extracts and displays the available metadata.

Example:

```bash
./spider -r -l 3 -p ./data/ https://example.com

./scorpion ../spider/data/*
```

---

## 🛠️ Technologies

- **Go (Golang)** – Main programming language.
- **Command-line interfaces** – Program execution and argument handling.
- **Web crawling** – Recursive website traversal and image discovery.
- **Image processing** – Image file parsing and metadata extraction.
- **EXIF** – Extraction of embedded image metadata.

---

## 🎯 Project Goals

This project focuses on developing practical programming skills in Go, including:

- Working with command-line arguments and flags.
- Making HTTP requests.
- Parsing HTML content.
- Implementing recursive algorithms.
- Handling files and directories.
- Processing different image formats.
- Extracting and interpreting metadata.
- Structuring Go applications into reusable packages.

---
