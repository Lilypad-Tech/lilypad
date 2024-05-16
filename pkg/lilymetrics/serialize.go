package lilymetrics

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FileType string

const (
	FileTypeFile      FileType = "file"
	FileTypeDirectory FileType = "directory"
)

type FileInfo struct {
	Name     string     `json:"name"`
	Path     string     `json:"path"`
	RelPath  string     `json:"relpath"`
	Size     int64      `json:"size"`
	Type     FileType   `json:"type"`
	Contents []FileInfo `json:"contents,omitempty"`
	Content  string     `json:"content,omitempty"`
	MIMEType string     `json:"mimeType,omitempty"`
}

func getMIMEType(fileType string) string {
	// Map common file extensions to MIME types
	mimeTypeMap := map[string]string{
		".mp3":  "audio/mpeg",
		".wav":  "audio/wav",
		".ogg":  "audio/ogg",
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		".gif":  "image/gif",
		".mp4":  "video/mp4",
		".avi":  "video/x-msvideo",
		".mkv":  "video/x-matroska",
		".gltf": "model/gltf+json",
		".glb":  "model/gltf-binary",
		".txt":  "text/plain",
	}

	mimeType := mimeTypeMap[fileType]
	if mimeType == "" {
		return "text/plain" // Default MIME type for unknown files
	}
	return mimeType
}

func serializeFolder(folderPath string) (FileInfo, error) {
	var folderInfo FileInfo
	folderInfo.Name = filepath.Base(folderPath)
	word := "downloaded-files"
	index := strings.Index(folderPath, word)

	if index == -1 {
		fmt.Printf("'%s' not found in the string\n", word)

	}
	folderInfo.RelPath = folderPath[index+len(word):]
	folderInfo.Path = folderPath

	folderInfo.Type = FileTypeDirectory

	// Open the folder
	folder, err := os.Open(folderPath)
	if err != nil {
		return folderInfo, err
	}
	defer folder.Close()

	// Read its contents
	files, err := folder.Readdir(-1)
	if err != nil {
		return folderInfo, err
	}

	for _, file := range files {
		var fileInfo FileInfo
		fileInfo.Name = file.Name()
		folderInfo.Name = filepath.Base(folderPath)
		word := "downloaded-files"
		index := strings.Index(folderPath, word)

		if index == -1 {
			fmt.Printf("'%s' not found in the string\n", word)

		}
		substring := folderPath[index+len(word):]
		//folderInfo.Path = substring
		fileInfo.RelPath = filepath.Join(substring, file.Name())
		fileInfo.Path = filepath.Join(folderPath, file.Name())
		fileInfo.Size = file.Size()

		if file.IsDir() {
			fileInfo.Type = FileTypeDirectory
			contents, err := serializeFolder(fileInfo.Path)
			if err != nil {
				return folderInfo, err
			}
			fileInfo.Contents = append(fileInfo.Contents, contents)
		} else {
			fileInfo.Type = FileTypeFile
			fileInfo.MIMEType = getMIMEType(filepath.Ext(fileInfo.Name))

			// Read and encode file content
			if fileInfo.MIMEType == "text/plain" {
				content, err := os.ReadFile(fileInfo.Path)
				if err != nil {
					return folderInfo, err
				}
				fileInfo.Content = string(content) //base64.StdEncoding.EncodeToString([]byte(content))
			}
		}

		folderInfo.Contents = append(folderInfo.Contents, fileInfo)
	}

	return folderInfo, nil
}

func FolderInfo(folderPath string) string {
	//folderPath := "/path/to/your/folder"
	//outputFile := "folder_structure.json"

	folderInfo, err := serializeFolder(folderPath)
	if err != nil {
		fmt.Println("Error:", err)

	}

	// Marshal folderInfo struct to JSON
	_, err = json.MarshalIndent(folderInfo, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling folderInfo:", err)

	}

	// Create a map to hold the top-level structure of the JSON output
	result := map[string]interface{}{
		"Type":    "result",
		"Details": folderInfo, // Directly assign folderInfo struct
	}

	// Marshal the map to JSON
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling result:", err)

	}
	return string(resultJSON)
	// err = os.WriteFile(outputFile, jsonBytes, 0644)
	// if err != nil {
	//     fmt.Println("Error writing to file:", err)
	//     return ""
	// }

	// fmt.Println("Folder structure serialized to", outputFile)
}
