package resourceprovider

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"

	"github.com/rs/zerolog/log"
)

func decompressAndExtract(data []byte, destFile string) error {
	// Decompress the gzip layer
	gzipReader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		log.Debug().
			Str("decompress", "NewReader").
			Msgf("create gzip reader failed: %v", err)
		return err
	}
	defer gzipReader.Close()

	// Extract the tar archive
	tarReader := tar.NewReader(gzipReader)
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of tar archive
		}
		if err != nil {
			log.Debug().
				Str("decompress", "Read").
				Msgf("read tar header failed: %v", err)
			return err
		}

		// Only extract the file with the specified name
		if header.Typeflag == tar.TypeReg { //} && filepath.Base(header.Name) == filepath.Base(destFile) {
			outFile, err := os.Create(destFile)
			if err != nil {
				log.Debug().
					Str("decompress", "Create").
					Msgf("create file failed: %v", err)
				return err
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				log.Debug().
					Str("decompress", "Copy").
					Msgf("write file failed: %v", err)
				return err
			}
			outFile.Close()
			if err := os.Chmod(destFile, os.FileMode(header.Mode)); err != nil {
				log.Debug().
					Str("decompress", "Chmod").
					Msgf("chmod file failed: %v", err)
				return err
			}
			break
		}
	}

	return nil
}
