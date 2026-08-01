# 🛠️ Git Workspace Status Tracker

A high-performance Terminal User Interface (TUI) tool built in Go that automatically scans Windows directories to track the status of local Git repositories. 

This project was built to explore backend developer tool workflows, Go filesystem cross-platform operations, and native system process execution.

## 🚀 Features
- **Automated Traversal:** Recursively crawls folders safely using Windows-optimized directory path handlers.
- **Native Execution:** Interacts directly with the local machine's `git` system process.
- **Interactive UI Layout:** Fully styled terminal navigation framework built with Charmbracelet Lip Gloss and Bubble Tea loops.
- **Desktop Actions:** Pressing `Enter` fires a native system execution to launch Windows File Explorer on the selected row.

## ⚙️ Built With
- **Language:** Go (Golang)
- **TUI Architecture:** [Bubble Tea](https://github.com)
- **Terminal Styling Stylesheet:** [Lip Gloss](https://github.com)

## 📦 Local Setup Instructions

### Prerequisites
Ensure you have Go and Git installed on your Windows environment path.

### Installation
```bash
# Clone the repository
git clone https://github.com
cd git-tracker

# Initialize dependencies and launch the executable app loop
go run main.go
```
