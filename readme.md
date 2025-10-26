# showmd

A fast and secure Markdown to HTML rendering tool built with Go. Uses [`gomarkdown/markdown`](https://github.com/gomarkdown/markdown) for parsing and [`bluemonday`](https://github.com/microcosm-cc/bluemonday) for secure HTML sanitization.

## Features

- 🚀 Fast Markdown to HTML conversion
- 🔒 Secure HTML output with sanitization
- 📦 Standalone binary - no dependencies needed
- 🌍 Cross-platform support (Windows, Linux, macOS)

## Quick Start

```bash
showmd input.md
md show on http://localhost:4000
```

---

## 📥 Installation

### 📦 Step 1: Go to Releases Page

1. Visit the [Releases page](https://github.com/your-username/go-markdown-render/releases)
2. Download the latest release for your operating system

### 🧩 Step 2: Download the correct binary

Download the appropriate file for your system:

| OS      | Architecture | File                       |
| ------- | ------------ | -------------------------- |
| Windows | amd64        | `showmd_windows_amd64.exe` |
| Linux   | arm64        | `showmd_linux_arm64`       |
| macOS   | arm64        | `showmd_darwin_arm64`      |

---

### 🧱 Step 3: Rename the binary

Rename the downloaded file to just **`showmd`** for consistency.

#### Windows:

```powershell
rename showmd_windows_amd64.exe showmd.exe
```

#### Linux / macOS:

```bash
mv showmd_linux_arm64 showmd
# or
mv showmd_darwin_arm64 showmd
```

---

### ⚙️ Step 4: Make it executable (Linux / macOS only)

```bash
chmod +x showmd
```

---

### 🌍 Step 5: Move it to a system PATH directory

#### **Windows:**

1. Move `showmd.exe` to a permanent location, e.g.
   `C:\Program Files\showmd\`
2. Add that folder to your system PATH:

   - Press **Win + S**, search **Environment Variables**
   - Click **Edit the system environment variables**
   - Click **Environment Variables…**
   - Under **System variables**, select **Path → Edit → New**
   - Add the path:

     ```
     C:\Program Files\showmd
     ```

   - Click **OK** to save everything.

#### **Linux / macOS:**

Move `showmd` to `/usr/local/bin` (or any PATH directory):

```bash
sudo mv showmd /usr/local/bin/
```

---

### ✅ Step 6: Verify installation

Run:

```bash
showmd --version
```

If it prints version info or help text, it's installed correctly.

---

## 🛠️ Building from Source

### Prerequisites

- Go 1.24.3 or later
- Node.js with pnpm (v10.12.2+)

### Build Steps

```bash
# Clone the repository
git clone <repository-url>
cd go-markdown-render

# Install Node.js dependencies
pnpm install

# Build the binary
node build.mjs
```

The compiled binary will be available in the output directory.

---

## 📖 Usage Examples

### Basic Conversion

```bash
# Convert markdown from stdin
echo "**Bold text** and *italic*" | showmd
```

### File Conversion

```bash
# Convert a markdown file
showmd README.md

# Save to HTML file
showmd README.md > README.html
```

### Pipe from other commands

```bash
# Convert curl output
curl https://raw.githubusercontent.com/user/repo/main/README.md | showmd
```

---

## 🔧 Technical Details

- **Language:** Go 1.24.3
- **Markdown Parser:** [gomarkdown/markdown](https://github.com/gomarkdown/markdown)
- **HTML Sanitizer:** [bluemonday](https://github.com/microcosm-cc/bluemonday)
- **Build System:** Node.js with zx and execa

---

## 📝 License

See LICENSE file for details.

---

## 🤝 Contributors

- Ned Palacios ([@nedpals](https://github.com/nedpals))
