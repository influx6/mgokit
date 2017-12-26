package static


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

type fileData struct{
  path string
  root string
  data []byte
}

var (
  assets = map[string][]string{
    
      ".tml": []string{  // all .tml assets.
        
          "mongo-api-backend.tml",
        
          "mongo-api-json.tml",
        
          "mongo-api-readme.tml",
        
          "mongo-api-test.tml",
        
          "mongo-api.tml",
        
          "mongo-solo-readme.tml",
        
          "mongo-solo.tml",
        
      },
    
  }

  assetFiles = map[string]fileData{
    
      
        "mongo-api-backend.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\x41\x8b\xa3\x40\x10\x85\xef\xfe\x8a\x3a\x26\x20\xfa\x17\x76\x8d\x6c\xd8\xcb\x26\xb0\xec\x69\xd9\x43\xdb\xbe\x68\x6f\xda\xee\xa6\xbb\xcc\x44\x82\xff\x7d\x30\x86\x44\x06\x1d\x48\xe6\x24\x56\xd7\xab\xf7\xbe\xa2\xd2\x94\x2e\x97\xe4\x37\xfb\x56\x72\xb2\x2b\xfe\x43\x72\xf2\x4b\x34\xe8\xfb\x3c\xcb\x84\x3c\xc2\x94\x54\xe2\xa0\x0c\x02\x09\x2a\x6e\x95\xb7\x5a\xc9\x9a\x3c\x9c\x47\x80\xe1\x40\x5c\x83\x2a\x75\x52\xa6\x8a\xd2\x94\x1a\x70\x6d\xcb\x40\x38\x3b\x1b\x50\x52\xd1\x5d\x1b\xf2\x8c\x54\xe3\x34\x1a\x18\x16\xac\xac\xa1\x83\xf5\x13\x29\x71\xe7\xb0\x14\x27\x19\x06\x7f\xbb\xeb\xa3\xcf\x7a\x1f\xd1\x95\x61\xf8\x83\x90\xb8\x44\x44\x1b\xdb\x1a\x5e\x49\x3e\x93\xb4\x86\x71\xe6\x64\x33\x7e\xd7\xb4\x52\x86\x63\x82\xf7\xd6\xaf\x23\xa2\x1c\x1a\x8c\xb9\xd6\x98\x5c\x5b\x68\x25\x7f\xe6\x14\xd8\x2b\x53\xad\x47\xd5\x30\xde\x43\x2c\x89\xa0\xd1\x4c\xd2\xee\x85\x3c\x8a\x6a\xa0\x5a\x20\xb8\x4d\xa5\x88\x68\x8b\xd9\xcc\x33\x41\x68\xf5\x84\x43\x4c\x0f\xdc\x3f\xae\x5c\x4c\xfe\xc1\xe5\x65\x94\x91\xe4\xbb\xd6\x59\xb7\xf3\x25\xfc\xbc\x9b\x1d\x9e\xee\x56\xd7\xbf\xac\x9b\x00\xfe\xfd\xf7\x22\xe2\x16\x9c\x75\x3f\x14\x74\x39\x6f\x7c\x44\x77\xb7\x3d\x09\xdd\x62\x72\x3a\xfd\x17\x56\x3b\x32\x3f\x0f\x1b\x93\x13\xd5\x35\x44\x4c\x1e\xc1\x59\x13\xb0\x87\xdf\xdf\x8a\xaf\xec\x62\x7a\xe2\x7d\xf4\x1e\x00\x00\xff\xff\x3b\x48\x8e\xe6\xf9\x03\x00\x00"),
          path: "mongo-api-backend.tml",
          root: "mongo-api-backend.tml",
        },
      
        "mongo-api-json.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x91\xcf\x4a\xc4\x30\x10\xc6\xcf\x0d\xe4\x1d\xc6\x3d\x48\x0b\x4b\xf6\xae\xf4\x05\x44\x76\xc5\xc5\x93\x08\x3b\xcd\x4e\xd7\xd6\x76\x22\x49\xea\x1f\x42\xde\x5d\x92\xae\xe0\x41\xd0\x5e\x12\x18\x92\xdf\xf7\xe3\x1b\x29\x36\x1b\xe8\x9d\x61\x68\xbb\x0f\x3f\x59\x72\xa0\x94\x92\xe2\x0d\x2d\x94\x52\x40\x08\x6a\xef\xed\xa4\xbd\xda\x35\x3d\x69\xaf\xb6\x38\x52\x3e\x62\xbc\xd9\xef\xb6\x50\xc3\x21\x04\x18\xf1\xf5\x1e\xf9\x68\xc6\x3c\x3b\x7f\x81\x55\xe3\x0c\xaf\x60\xd5\xe7\x2b\xc6\x83\x14\x95\x14\x39\xf2\xd6\xe0\xf1\x4f\xb6\x25\x3f\x59\x76\x80\xc0\xf4\x0e\x1d\x3b\x8f\xac\x09\x4c\x0b\xf8\x43\xec\x0e\xf5\x0b\x9e\x28\x46\xf5\x2b\x30\x46\x25\x45\x3b\xb1\xfe\x57\x66\xa9\x0d\x7b\x62\x0f\xce\xdb\x8e\x4f\x15\x94\x0b\x82\xd6\x40\xd6\x1a\x5b\x41\x90\xa2\x48\x0d\xd2\x40\xe3\x12\xd3\x54\x4e\xd1\xb5\x09\x03\x57\x75\xde\x8b\x7a\xe0\x11\xad\x7b\xc6\xa1\x7c\x7c\x6a\x3e\x3d\x7d\x1b\x56\x6b\xb8\x4c\xfc\xea\x3a\x3f\xbf\xa8\x81\xbb\x21\x27\x17\x73\x6f\x4b\x82\xc3\xec\x2e\x45\x31\x3b\x9c\x09\x89\xbf\x4e\x5c\x29\xf2\xfc\x2b\x00\x00\xff\xff\x39\x66\x83\xa1\x2e\x02\x00\x00"),
          path: "mongo-api-json.tml",
          root: "mongo-api-json.tml",
        },
      
        "mongo-api-readme.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\x41\x8f\xd3\x30\x10\x85\xef\x95\xf2\x1f\x06\xf5\x92\xa2\xe2\xde\x57\xe2\x50\x52\x54\xed\x01\xa8\x16\xf6\x54\xad\x14\xd7\x99\x3a\x66\x1d\x4f\x64\x4f\x44\xa4\x2a\xff\x1d\x39\xe9\xd2\x16\x52\x28\x82\x5c\x9c\x19\x65\xde\x7c\xef\x39\x87\x83\xf8\xcc\xbe\x51\x2c\x3e\xed\xbe\xa2\x62\xf1\x51\x56\xd8\x75\xf0\x81\x9c\xa6\xd5\x3b\x58\x6e\xee\x93\xc9\xdb\x3f\x3f\xc9\x64\xfb\x6a\xbb\x26\x78\xc0\x9a\x3c\x43\x26\x7d\xf1\x94\x96\xcc\x75\xb8\x5b\x2c\x34\xf9\xbe\xad\xa4\x2f\x84\xa2\x6a\xb1\x93\x85\xc6\xc5\xe1\x20\x36\x52\x3d\x4b\x8d\x1b\xc9\x65\xd7\xcd\x7e\x33\x31\x94\xbf\x8e\x24\x93\x64\x72\x83\x07\x30\x01\x24\xc8\x86\xe9\x8d\x46\x87\x5e\x32\x16\x90\x3d\x3c\xae\xc0\x54\xb5\xc5\x0a\x1d\x4b\x36\xe4\x60\x4f\x1e\xb8\x44\xc8\x47\x45\x8f\xca\x39\x18\x07\xf5\xc0\xd1\x7f\xb9\x79\xd6\x62\x00\xca\x45\x24\xfa\x52\x22\xec\xc9\x5a\xfa\x66\x9c\x86\x0a\xb9\xa4\x02\xb0\x35\x81\x43\xbf\x41\x35\x81\xa9\x02\xaa\x23\x89\x21\x17\xee\xe2\xd4\x74\x0a\xef\x5b\x54\xf1\x35\xcf\x73\x4d\xc9\x24\x96\xa9\xe2\x16\x14\x39\xc6\x96\x45\x36\x9c\x73\xd8\xb7\xb0\x6f\x9c\x4a\x15\x59\x78\x5d\x69\x12\x19\x59\x8b\x2a\x8a\xcd\x00\xbd\x27\x7f\x3c\x7a\xad\x6b\x4c\xe1\x05\xca\xb8\xde\xf5\x29\x9b\x98\x99\x0c\x50\xa3\x67\x69\x5c\x9c\x60\xea\x03\x7b\x21\xcd\xa8\x71\x7c\x86\xda\xd7\x63\xac\x33\x48\x8d\xe3\xf9\x11\xea\x07\x4e\x94\xf0\x28\x19\xcf\x35\xfa\xc6\xb8\x61\xb4\x58\xc1\xe9\x52\x8e\x7f\x41\xd7\x89\x2b\xb7\xff\xb3\xfd\xe9\x14\xd6\x78\x0e\xbc\xc6\x51\xdc\x39\xd4\xcd\xce\x1a\x75\xbf\x82\xc0\xde\x38\x3d\x83\xf4\x2f\xd6\x8e\xf9\x5c\x23\xc3\xd2\xda\xcb\xdd\x4b\x6b\xaf\xa4\xb5\x7d\xfa\xc7\x7d\x8f\x75\x71\x99\xeb\xd0\xb8\xc9\xed\x7f\x09\x7a\x85\x16\x2f\x00\x86\xc6\x8d\x71\x9f\xe4\xbe\x07\x00\x00\xff\xff\x64\x58\x5b\x5b\x9d\x04\x00\x00"),
          path: "mongo-api-readme.tml",
          root: "mongo-api-readme.tml",
        },
      
        "mongo-api-test.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x59\xdb\x6e\xe3\x36\x10\x7d\xb6\x01\xff\xc3\x54\x40\x01\x79\xeb\x72\x37\xfb\xe8\xc2\x0f\x71\xbc\x4d\x5a\x34\x71\xb0\x72\xb0\x8f\x06\x43\x8e\x1c\x36\x94\x98\x92\x23\x27\x81\xa1\x7f\x2f\x74\xf1\x6d\x1d\x27\x96\x73\x69\xb1\x51\x80\x18\x10\xc5\x99\x33\x67\x78\xe6\x18\x96\xa6\xdc\x82\xdf\x6a\x02\x00\xe0\x14\x63\x72\xd0\x83\x08\xc9\x2a\xe1\xd8\x19\xde\xfa\x22\x71\x64\x22\x16\x10\x17\xd7\x03\xe5\x6e\x34\xbf\xf7\x8d\x63\x01\x49\x93\x50\xbb\xdd\x6a\x16\xb1\xc2\xc4\xa1\x9a\x64\xb1\xf2\x92\x1d\xe5\x17\xb3\xe2\x4e\xf6\x77\x6a\x24\x76\x21\x9a\x18\x76\x6a\x62\x43\x26\x56\xa2\xb3\xbc\x3b\xe8\x77\xc1\x38\x76\x8c\x84\xf1\xd4\xf7\x66\x33\x16\x90\x4d\x04\xb1\x73\x2e\xae\xf9\x04\xd3\x74\x7c\x3a\x3c\x3b\x1e\x8e\x47\x5f\x82\xd1\x78\xd0\xf7\xda\x2b\xc1\x27\xc6\x51\x95\xf0\x93\x61\x30\x5a\x4b\x70\xe1\xd0\x56\x49\x70\x11\x7c\xf9\xba\x96\xe0\x30\xa1\xab\x6a\x14\x0e\x2f\x46\x27\xdf\xd1\x38\xe7\xce\xdd\x1a\x2b\xab\xa4\x39\x3f\x0c\x82\x6f\xc3\xaf\x83\x45\xa2\x74\x7e\x1a\x84\x8e\x8e\x8c\x86\x1e\x78\xb3\x99\x36\xb7\x68\x61\x9e\x69\x78\xf9\x37\x0a\x62\x67\x3c\xc2\xfc\x23\x4d\xc7\xd9\xee\xb1\x30\x5a\xa3\x20\x65\x62\xaf\xd5\xcc\x8f\xf5\xe3\x47\x18\xa1\xa3\x63\xa4\x65\x1d\x2b\xd1\x69\x0a\x53\xae\x95\xe4\x84\x0e\xe8\x0a\xc1\x66\xaa\xc1\x29\xd7\x60\x42\xe0\xb0\x25\x28\xcf\x6b\x51\x18\x2b\x21\xb4\x26\x02\x0e\x91\x89\x27\x46\x5e\xb2\x56\x33\x4c\x62\xf1\x04\xa8\x4f\xf0\x21\x2b\x58\xc5\x13\x36\x6a\x97\x1a\xe3\x37\x0a\xba\x85\xf6\x32\xcd\x96\xf4\x3b\xa5\xa2\x3b\xf3\x1b\xa7\x19\xd0\xa0\xef\x17\x62\x5d\x11\x2f\xdd\x75\x40\xf0\x58\xa0\xce\xd2\x08\x13\x13\xde\x11\xfb\xa6\xe8\x6a\xa4\x22\x34\x09\xf9\xf3\xb5\x3e\x17\xd7\x13\x6b\x92\x58\xfa\xed\x0e\x1c\x7c\x82\x0f\x40\x2a\x42\x16\xa0\x30\xb1\x6c\x17\xf9\x24\x86\x68\xcb\x84\xfe\x02\x05\x35\x46\x1d\x40\x6b\x33\x8c\x50\xdd\x51\x62\xd1\xb1\xbf\x0c\x97\x0f\x52\x2d\xf9\xfe\x19\x0c\xcf\xfc\xc5\xee\xa7\x76\x96\x05\xa8\x30\xc7\xf9\xa9\x07\xb1\xd2\xb0\x32\x88\x59\x67\x1c\xfb\x9d\x2b\x8d\xd2\xf7\x82\x44\x08\x74\x2e\x4c\xb4\xbe\x07\x6d\xb8\x44\x09\x59\x16\x08\x8d\xdd\x76\x7e\xe5\xd9\x75\xe1\xe7\x5f\xfe\x61\x5e\xce\xa7\xbd\x50\xdf\x12\x22\xd3\xf3\x33\x21\xbc\x45\xe7\x8a\x7e\xf2\x1b\xc5\x06\xa8\x91\xd0\xcf\x0f\x2c\xeb\x27\x3b\x4f\x2e\xb5\x12\x7f\x0c\x16\x7b\x4b\xea\xdd\x5e\xbe\xff\xc8\x22\x5f\xdd\xdf\xfe\xad\x7a\x63\xb8\xcc\x8a\x9e\x4b\xf6\x91\xb2\x55\x4c\x06\xe4\xe5\x5e\xad\xa9\x0a\xc2\x96\xdd\x19\x17\xa2\x2a\x08\x1f\x23\x3d\xdc\x9d\x7d\x65\x51\x4e\x35\x4a\x70\x64\xec\x6e\x35\xe6\x73\xbd\x67\x23\x9e\x81\x97\xf7\x24\x5d\xb7\xae\x43\xad\xf7\x71\x2f\xad\x9f\xed\x5f\xdb\x91\xff\xef\x16\xd6\xd8\xf4\xaf\xc6\x1b\x99\x57\x63\x43\xa2\x8d\xc6\x6b\x79\x56\x23\x6d\x35\x1b\xb5\x5b\xbd\xb1\x5b\x15\x41\xae\x33\xb7\xad\xee\xc2\xb7\x0e\xb5\x2e\xa8\x7b\xdc\x09\xaf\x03\xde\x4d\xde\xad\xb1\x92\x5e\x07\x7e\x3d\xc8\xfe\x5f\xc4\xc7\xb2\xe1\x2e\xab\x78\x03\x17\xab\x88\xb6\xd2\x29\x15\x82\xc6\xd8\x2f\x83\xdb\xd0\xeb\xc1\xa7\xea\x64\x49\x23\x77\x04\x07\x55\x7d\xb4\x32\xcf\x7d\x81\x76\x36\xec\xa1\x95\x68\xfb\xf7\xff\x9d\x6f\xf7\xef\xf3\x12\x6a\xfb\xae\xed\xbb\xb6\xef\x0d\xef\x9e\x4f\xc7\x16\x0b\xaf\xad\xfb\xc7\xb6\xee\x2d\xfb\x8b\x99\xf8\xce\xb3\x45\xb6\xa8\x4c\xbc\xf3\x83\x82\x5b\x45\x57\x0f\x1b\xf6\xa3\xb0\xb5\x53\xd7\x4e\xfd\x0e\x9d\x7a\x87\xb1\xbc\xb8\x91\x9b\x63\x99\x14\x8b\xaf\x36\x94\x05\x68\x3d\x94\xf5\x50\xbe\xc3\xa1\x5c\x3e\x03\xfe\x5c\x2b\xe9\x49\x25\xe5\x13\xf7\x79\xa1\x18\xe8\xad\x2b\xe8\x61\x01\x95\xfe\xb2\x10\xd0\x32\xbe\xbc\xde\x47\x51\x85\x29\xbe\xba\xa6\xaa\xc3\xec\x6a\xf5\xc5\x18\x6e\xfc\x6a\x8e\xcc\xcb\xbc\xa9\x79\x14\xb5\xf6\xfa\x1f\x76\x42\xdf\xa7\x7f\xaf\x33\xde\xfe\x0d\xb7\x07\xf5\x7c\x24\x5f\x9f\x7c\x75\x98\x75\xfa\xbb\xbc\x6d\x2a\xd8\xf7\xaa\xb0\x0f\xf3\x45\x20\x03\x13\x24\x90\x79\x63\x2b\x17\xb9\x5b\x07\x5e\x04\x2a\x6d\x35\xff\x0d\x00\x00\xff\xff\x33\xcc\xcb\x04\x3a\x21\x00\x00"),
          path: "mongo-api-test.tml",
          root: "mongo-api-test.tml",
        },
      
        "mongo-api.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\xd9\x6f\xdb\x46\xb7\x7f\xb6\x01\xff\x0f\xe7\x13\x3e\x04\x64\xaa\x32\xe9\x7d\xb8\x0f\x6a\x5d\xa0\xb6\x93\xde\x00\xcd\x72\xe3\xf4\xf6\x21\x08\x82\x11\x79\x24\xb3\x21\x67\x74\x39\x23\x2f\x10\xf4\xbf\x7f\x38\xb3\x71\x11\xb5\x9a\xb1\x65\xd7\x7d\xa8\xa3\xe1\x2c\x67\x3f\xbf\x33\x33\x94\x5e\xbc\x00\x2c\x0a\x51\x48\x88\xa2\xe8\xe8\xf0\x92\x15\x10\x1c\x1d\x02\x00\xbc\x2a\x8a\x77\x42\xbd\x16\x53\x9e\xc0\xb1\xed\x14\xbd\xc3\xab\xa0\x57\x60\x2c\x8a\x04\xb8\x50\x30\xa2\xc7\xbd\xf0\xe8\x30\x3c\x3a\x3c\x3a\x7c\xf1\x02\x4e\x05\x1f\xa5\x63\x98\x14\xe2\x32\x4d\x50\x42\xac\x3f\x4f\x0b\xa6\x52\xc1\x61\x24\x0a\x6a\xe1\x18\xab\x94\x8f\x41\x09\x60\x90\x0c\xa3\xa3\x43\x75\x33\x41\x37\x56\xaa\x62\x1a\x2b\x98\x1d\x1d\x1e\xfc\x8f\x90\x8a\x68\xa1\xb6\x94\x8f\x8f\x0e\x0f\x7e\x9b\xaa\x8b\xb3\x93\x6a\x8b\xfe\x54\xeb\xf3\xa7\xc4\xa2\xde\xf2\x81\x49\x79\x45\x34\xfb\x96\xb7\x22\x41\xdd\x27\x1f\x8b\x88\x3e\x1c\x1d\xce\x2d\x0b\xbf\xa3\x3a\x47\x29\x89\x5e\xa6\x14\xe6\x13\x25\x89\xd2\x02\x55\x91\xe2\x25\x82\xba\x40\x18\xa7\x97\xc4\x80\xb4\xfd\x88\x2f\xdb\x8c\xdc\xf2\x1c\x1d\x1d\x8e\xa6\x3c\xae\x4c\x17\x98\x07\x96\xcf\x10\x82\xe7\xb4\xb8\x7d\xd8\x37\x22\x0e\x35\xdf\x29\x1f\x09\x18\x1c\x6b\xe2\xce\x52\x96\xbd\xe1\x23\x41\xed\x07\xbf\x25\x49\x21\x07\x44\xf7\xe7\x2f\x86\x97\x99\x5d\x8d\x44\x35\xef\x53\x9f\x4f\x69\x8e\x62\xaa\x06\x00\xff\xfd\x12\x9e\x83\x4a\x73\x8c\xce\x31\x16\x3c\xd1\x8f\xcf\x98\x62\x43\x26\x71\xe0\xe8\x34\x22\xd5\xcf\x48\x72\x9c\xe5\xe5\x33\x6a\xd0\x4f\x9c\x04\xfd\x13\xd7\x40\x4f\xb5\xe4\x0e\x48\xfb\x05\x32\x85\xc0\xbc\x60\xae\x2e\xd2\xf8\x02\x72\x96\x72\xc5\x52\x2e\x81\xc1\x44\x88\x0c\xc4\x08\xa4\x88\xbf\xa1\xf2\xd6\x20\xb8\x34\x53\x28\x01\x62\x5a\xc0\x5b\xc1\xc7\xe2\xec\x24\x3a\x3a\x3c\x90\x28\xb5\x70\xaa\x02\xf9\x2b\x55\x17\x24\x94\xe0\x19\x89\x2a\x24\x91\x8d\x74\x9f\x7f\x1d\x03\x4f\x33\x2d\xc4\x83\x02\xd5\xb4\xe0\xf4\x59\x8f\xf7\x84\xa6\x23\xc7\x84\xb6\x82\xe3\x63\x78\x69\x06\xd4\x5a\xad\x65\x70\xa1\x04\x4f\x63\x3f\x58\xa2\x8c\xce\x51\x51\x9f\xa0\xd2\xbf\x0f\xaa\x98\xa2\xf6\x02\xb7\xae\xa6\x9b\xa7\x59\x69\x59\xef\xf0\xca\xf2\x05\xa6\x0f\x09\x84\xe3\x15\xa4\x5c\x2a\xc6\x63\x24\xc1\x30\xc3\xfb\x39\x16\x97\x58\x38\x23\x2a\x47\x36\x8d\xc8\x4d\x48\x0c\x90\xf7\xe6\x1c\xf2\x72\xfc\xd1\xe1\x41\xce\x23\xeb\x58\xc7\x96\xed\x92\xc4\x67\x39\x2f\xa9\xab\x0c\x83\x04\x47\x29\x47\x22\x4f\xb7\x56\xd4\x04\x39\xe3\x6c\x8c\x64\xef\x4c\xc1\x70\x9a\x66\x89\xd4\xc3\x59\x96\x89\x2b\x09\x53\xc9\xc6\x96\x0f\xeb\x24\xf5\x10\xa0\x04\x8c\x91\x63\x41\x76\x42\xac\xeb\xf9\xf5\x04\xd6\x66\x24\x30\x9e\x40\x62\xad\xd4\x8b\x46\xba\x28\x51\x25\xb3\x12\x2a\x4e\x2d\x6b\x15\x59\x37\x84\xec\x6c\xb2\x36\xfd\xa8\x10\x79\xd5\xa3\x6b\xc4\x3a\xe9\x07\x39\x3c\xaf\x2c\x1b\xd2\xe4\x81\xf3\x5f\xe7\x4f\x7d\x58\xea\xce\x55\x0b\xae\x84\x83\xdc\x2a\x66\x03\xf3\x5d\xb0\xe1\xd2\xc6\xa2\xb3\x13\x3f\x53\x74\x76\x12\xf6\x5b\x0c\x6f\x36\x8b\xce\xb5\xa8\xa2\xf7\xc3\xbf\x31\x56\xd1\x3b\x96\xe3\x7c\xfe\x3a\xc5\x2c\x91\xa5\xb2\x39\xa4\x5c\x61\x31\x62\x31\x5a\xcf\xc5\xeb\x89\x90\x28\x21\x47\x75\x21\x12\x1b\x06\x69\x61\x06\x39\x9b\x68\x35\x67\x99\x51\xbf\x52\x45\x3a\x9c\x2a\x9a\x47\x4a\x11\xa7\x4c\x61\x02\x57\xa9\xba\xd0\xe2\x35\x6b\x24\x56\x63\xd3\x02\x81\xd1\xc2\x71\x9a\x60\x02\xc3\x1b\xdd\xc7\x3f\x73\xaa\x5e\x4d\x76\x85\x58\x92\x97\x69\x25\xb5\xe4\x6c\xf2\xd9\x04\xc7\x2f\xbe\xcb\x6c\xee\x34\xb2\x56\x2a\xa7\x82\xcb\x69\x5e\x75\x82\x45\xb9\xb0\x38\x46\xca\x0b\x5e\x0c\x64\x50\xf6\xd9\x55\x9a\x65\x30\x44\x32\x25\x9a\x27\xd1\x6b\xa5\x5c\x89\xaa\x9d\xa5\xf9\x24\xc3\x1c\xb9\xce\x83\x5d\x08\xc5\x53\x5d\x97\x8a\x6d\x5e\x22\x93\xd0\xc8\xa4\x14\x89\x8b\x26\xa5\xff\x2f\xb3\x08\xe6\x6c\x82\x72\x9f\xcd\x8d\xc4\x0b\x2b\x63\x89\x77\x0d\xed\x72\xa6\xc9\xda\xbe\x63\xc6\xad\x57\x27\x7a\x73\xf7\x5a\xab\xcc\x1a\x2f\xa5\x9c\x0d\x2f\x16\xa5\xc0\xd9\x09\x9c\x7e\xfc\xf3\x0c\xc4\x04\x8d\xe3\x9b\x88\x36\x95\xc4\x90\x09\x80\x4c\x6a\x6d\x4c\x79\x82\x45\x96\x72\xac\xc0\x96\xe5\x2b\x9b\xf5\x66\x84\xa7\x62\x91\x79\xf0\x01\x90\x0c\x1d\xe7\xf4\x29\x27\xe9\xc5\xd2\xfd\x8d\xde\x9a\xbf\xf4\x08\xb9\x9c\x16\x98\xbc\xe1\x09\x5e\xc3\x50\x88\x8c\x1a\x53\x1e\x0b\x32\x1f\x85\xcd\xf6\x04\xaf\x51\xc2\xe7\x2f\x24\x29\xfd\x6c\x55\x3c\xac\x26\x9d\xa5\x3c\x54\x12\x50\x50\xf2\xd0\x87\xbc\x49\x6d\x1f\x20\x17\x8e\xab\xbe\xa7\x25\x8a\x22\x4f\x4c\x08\xcf\x97\xae\x33\x33\xa0\xd3\xe5\xa5\x75\xfd\xe8\xbf\x64\x38\x80\x5c\xf4\xcb\x86\x58\x64\x04\x4f\xb2\x4a\x93\x25\x72\x00\x79\xa5\xd1\xd2\x36\x70\xff\xb0\x8f\xe6\xa5\xb0\x8c\xd8\x8d\x74\xab\x18\xd0\xb4\x53\xcc\x73\xb6\x93\x78\x4e\xbd\x87\xcb\x09\xc6\xe9\x28\x8d\x89\x94\xcc\x24\xcd\x32\x91\x24\xc3\x15\x42\x08\xab\x0b\x07\xd6\x39\xc1\xb2\x9c\x8e\x20\x4f\x86\x51\xcd\x22\x2a\xd2\x28\xb3\x85\xe7\x86\xfe\x98\xff\xe7\x1a\x3b\x59\x85\xbd\xc3\xab\x4f\x05\x8b\x31\xe8\x2d\xd7\x7a\x85\x0e\x42\xf7\x5a\xdc\x38\xc2\x42\x93\xe0\x26\x7a\x95\xa7\x2a\x70\x1f\x34\x1a\xdb\x74\xc6\xbe\x1b\x45\x38\xce\x10\x93\x47\xaf\x78\x12\x84\xa1\x86\x50\x96\xdf\x0c\x39\x49\x2c\xb2\x22\x0e\x3d\x54\x5b\xc1\xb4\xf9\xc7\x8b\x17\xf0\x66\x04\x57\x08\x17\x2c\xa1\xf8\x6d\x24\x39\xc4\x91\x28\xd0\x68\x0c\xae\x18\x55\x27\xc6\x8f\xfa\xa4\x38\x0e\xf2\x5b\x3a\xe9\xd3\xa8\x98\x71\x45\xf5\x8d\x9f\x4c\x2a\x31\xd1\x6a\x17\x13\x09\x43\x8c\xd9\x54\x6a\xbf\x19\xb1\x34\x73\x36\x10\x79\xba\xff\xb5\xa0\xa8\x67\xcf\xc0\x30\x52\xf7\xdc\x4d\x58\x49\x7c\x00\x94\x95\xd8\xa7\x15\x9a\x0c\xa3\x64\xa8\x6b\xb2\xd0\xaf\x5d\x47\x11\xde\x0f\x96\xa9\xed\x15\x09\x66\x14\xf4\x5e\x1b\x46\x94\x80\xd8\x60\xf8\x6a\x69\x93\xb6\x68\x2d\xe8\x95\xe6\xdd\xeb\xeb\x05\x62\x91\x35\xfb\x68\xb9\xf7\x34\xc5\x66\x29\xa3\xe1\x06\xd3\x1a\xd8\xd4\x99\xd6\xc6\x66\x69\x88\x4e\x33\x21\x31\xf0\x96\x51\x2e\x4c\x52\x70\xf2\x89\x4e\x03\x47\x84\xeb\x48\xb4\x7f\xb5\xb1\x88\xba\x16\x8c\x8f\x11\x2a\x16\x55\x15\x91\x95\xdd\xe0\xb8\xea\xb7\xaf\x2a\xfe\xa8\xc7\x84\x3f\x2f\x91\xf0\x96\x52\xb6\x91\xc4\x49\x79\x77\x09\x9b\x91\x96\xc9\x0d\xc5\xbf\x48\x75\xd3\x30\x8f\x75\x39\x53\xef\xd7\x54\x56\x4d\x61\x2b\xd9\x37\xb1\xe1\x7c\x1a\xc7\x88\x26\x64\x1a\xfe\x0d\xe0\xf6\xca\xec\x4a\x08\x61\xc3\x98\x16\xbc\xd1\x71\x57\x3e\x5e\x41\xf6\xeb\x94\xa7\xf2\x02\x13\x60\x49\xa2\x91\xdb\xe6\x54\x86\xb5\xa4\x56\x43\xe4\xa7\x62\xca\x55\x73\x7f\x81\x7a\x51\x06\x51\x42\xb1\x0c\xf8\x34\x1f\x62\x41\x51\xc6\x6e\xb8\xf8\x5a\x45\x63\x8f\xcd\x12\x8a\x5e\x27\x88\xd5\x35\xc1\x51\x85\xd7\x8a\x2a\x05\xfa\x1b\x42\x90\x72\x55\x2d\x53\x76\x4a\x14\x7a\xfe\x8e\x52\x84\x9d\x6b\xc3\xe4\x90\x4a\xcb\xc9\xab\xeb\x49\x5a\x60\x42\x4c\x86\x55\x8f\xb4\xde\x3c\xca\x95\xf7\x3f\x3b\x02\x2e\x98\x24\x20\x4b\xc3\x7a\xe1\x46\x26\xbc\xe8\xc1\x63\x54\x4e\x31\x71\x0b\xe1\x5d\xc6\xc7\x1f\x7f\xea\xb7\xc4\xc8\x32\x60\x95\x06\x6e\x91\xc3\xb2\x20\xb5\x05\x7b\x6c\x32\xc9\x6e\xba\x8f\xfc\x1b\xf2\x76\xd7\x49\xef\x7b\x2a\x73\x53\x96\x57\xa5\xbc\xff\x9f\x62\x71\x43\xec\x0f\xa5\xe0\xd1\xdb\x99\x45\x77\x3a\x50\x78\xd1\xb4\x64\xc2\xe8\x75\xca\x93\x40\x8f\x0e\x8d\x7f\x05\xe1\xcf\x7b\x26\x36\x4d\x5d\xaf\x6f\x78\xec\x56\xa6\x6b\x42\xd1\x19\x52\xc6\x4b\x2c\x0b\x1d\x10\xef\x29\x73\xc1\xdc\xeb\xa7\x8c\xfc\x66\xd1\x46\xe8\xcf\x85\xdd\x58\x5e\x0c\xf5\xb6\x0c\xa5\x0f\xbe\xe6\x98\x4c\x87\x59\x1a\xbf\x39\x8b\xf4\x8c\x1f\xf5\x18\xe9\x3b\xa6\x92\x2a\xda\x7c\x2a\x29\xd0\x5d\x22\x30\xdb\x1f\xd2\x04\x2e\x59\x36\xc5\x3e\x05\xbf\x02\xa5\xc4\x04\x30\x55\x17\x58\xc0\xf0\x06\x98\x36\x2e\x10\x05\xfc\x4d\x7f\x15\x1b\xeb\xd9\x05\xaf\x6c\x6d\x97\xc1\xfb\x03\x8b\xbf\xb1\x31\xce\xe7\xd1\x92\x80\x6e\xab\xdf\x8d\x33\x95\x91\x4b\x5b\xaa\xea\x7b\x7e\x6d\xf9\xd9\xa8\x8c\xb6\x4e\x5a\x66\xa9\x8e\xb2\x96\x9b\xac\x61\x18\x8e\xe4\x5e\x49\xfd\x3d\x64\xb6\x1d\x7c\x3a\x31\xf6\xb9\xc4\x27\xd6\xb2\xf5\x7d\xeb\x82\x07\x99\xf3\xd6\x72\x75\xd7\xd9\x6e\xaf\x55\xbc\x55\x1e\x2c\xe7\x2b\xc9\x1e\x78\xb2\xfb\x4b\xad\xa7\x2d\x55\x7e\xd4\x51\xd8\x26\xcb\x0e\xac\x69\xb5\x94\x6f\x9d\x19\x37\x50\xd3\x1a\x15\x58\x71\x1c\x9b\xa3\xae\xea\x79\xef\xac\xb5\xf2\xab\xf4\xa8\x55\x80\x6b\x55\x7a\xd7\x69\x78\x03\x49\x35\x33\x75\xbd\x38\xb3\xc7\x98\x95\x14\xcd\x92\xa4\x9a\x9f\xfd\x66\x5f\x7b\x7e\xae\x6e\xad\xaa\x0b\x6c\x6c\x50\xaf\x4d\x9d\xf7\x98\xd6\x6f\x95\xc2\x8d\xdc\xda\x53\x38\x66\x98\x6f\x23\x83\xdb\xe6\x78\x43\x4b\x57\x95\xa9\x9d\x6c\xb9\x61\x11\x7b\xd1\x87\x07\x95\xe8\xed\x46\xdf\xfa\x2c\xb0\x82\xb7\xa7\x6c\xbf\xff\xd9\xbe\xbe\xa1\xbb\x97\x8a\x5e\x99\xf3\x67\x33\x12\x43\x40\x16\xff\x9a\xc2\x90\x75\x52\xe8\x99\x13\xdf\x1e\x40\x08\xf3\x4a\x1a\x1a\xe9\x66\x2f\x54\xcd\x94\x3b\x1c\x5e\x48\x7e\x1b\x6e\xe4\xd6\x1f\x43\x79\xb2\xb4\xa4\x28\xf6\x87\xd1\x23\x0a\x61\x4b\xa2\xab\xf7\xbc\xe5\xb3\xaf\x12\xfd\xba\x51\xc4\xb8\x55\xea\x06\x9d\x5b\x34\xd7\x18\x14\x6e\xb9\x1d\xbc\x1a\x6d\xbd\xe1\x12\x0b\x15\x18\x1c\x17\x18\x9d\x85\x5d\xed\xae\x5b\x93\x5f\x2b\xf8\x1d\x2c\xbc\x2a\xd4\xad\x8c\x7f\x23\x99\xad\xc9\x50\xa7\x2b\x43\xf6\xb6\xf4\x87\xce\xbf\x30\x93\x58\xf5\xa0\x06\xca\x0e\x66\x33\x7d\xcd\xc1\x3b\x9e\x9e\x04\x7a\xf4\xb4\x07\xbd\xbf\xf5\x9f\xf9\x3c\xdc\x5a\xf9\xab\xa1\xf6\xde\xe8\x7c\x97\x0d\xaa\xfd\xd2\xfa\xc2\x2e\x95\xd5\x3b\x4f\xe6\xf3\x55\x40\xf8\x77\x54\xbf\x65\x99\xbf\xfa\x28\xf5\x11\x68\x61\x71\x69\x75\x8f\x8a\xf1\xa4\x72\xbb\x40\x66\x69\xf3\x5a\xc1\xfa\xed\x22\x75\x33\xc1\x87\x0a\x7c\x8d\x9c\xda\x81\xaf\x28\x12\x73\x53\x4d\xdf\x9b\xd0\x9f\x4e\x6e\xfc\xe7\x09\x1b\x23\xe8\x83\x98\x02\xe5\x44\x70\x89\x1f\xb0\xf8\x60\x1b\x43\x80\xe0\xf3\x97\x2d\x84\xd8\x07\xe8\xe2\x50\xc7\xb0\xd3\x11\x76\x76\x93\xad\xc3\xc4\x07\xf2\x2a\x55\xf1\x85\x95\x8c\x8c\x3e\x89\x3f\xc4\x15\x16\x81\x96\x98\xb9\x47\x17\x33\x89\xd0\x4b\x64\xdc\xeb\x43\x2f\x41\x19\xf7\x06\xa5\x13\x39\xc9\x1e\x43\xef\xc7\x1e\xfc\xe0\x3e\xbb\x1b\x73\xb0\x0f\x90\xdb\xdf\x21\xbe\x4d\x04\xdf\x08\x63\xe9\x0b\x83\x8f\xe8\xdc\x68\x33\xf6\x0e\xd2\x91\xf1\xa8\x5f\x8e\xe1\x25\x3c\x7b\xb6\xe0\x54\xbf\xf8\xbb\xbe\x36\x8a\xd5\xb0\xb7\x31\xd5\x93\x9b\xf7\x64\x3a\x64\x19\xd6\x5f\xbd\xdb\x86\x95\x7b\x99\x7e\x82\x0c\x79\x60\x3f\x84\xf5\x4b\x9a\x26\x82\x2e\x39\xce\x95\xd1\xd1\xe1\x81\x7e\xf4\xb1\x85\x14\x7f\x6e\xbb\xc1\xe5\x50\x2f\x07\xbb\xec\x25\x2b\xcc\x9a\x7f\x31\xae\x30\xb1\xc7\xe2\x9f\xc4\xb9\x62\x85\xa2\x00\xd1\x14\xd5\x4f\x6d\xa2\xfa\xd5\x49\xaa\x32\x15\x1c\x37\xbb\x51\x87\xda\xf4\xc7\xf0\x92\x08\x01\xc2\x14\x1b\x8c\x87\xe7\x9a\x8a\x96\x69\xaa\xc3\x5e\xc0\x7f\x69\x9a\x3d\xd1\xbf\xc2\x4f\x66\xf2\xda\xa8\x1f\x7e\xa0\xa6\xb9\x17\xc4\x0a\x18\xdf\xd8\x81\x3a\x19\xfc\x2f\xa5\xc6\x81\xb1\x00\x4b\x5b\x2f\xec\xc3\xe2\x08\x32\x53\x0b\xef\x5d\x93\xfe\xd8\x40\x2e\x3d\x49\x14\xa5\x7c\xfc\xd5\xf8\xc2\xc0\x01\xa3\x0a\xbd\x0d\x80\xdd\xd3\x2c\x7f\xb5\xe6\xf1\xf5\x4a\xf3\xde\x1b\xd4\x74\xd9\x18\xa1\xed\xd2\xcf\x0d\xb5\x58\xd8\xda\xf7\xe4\xa6\xd9\xdb\x36\x37\x7b\x93\x98\x17\x27\xd6\xd2\x6f\x76\x6d\xa8\xd4\x8d\x6a\x34\x57\x46\xcd\x5d\x69\x11\xee\x69\x4d\x7c\x27\x71\x78\xdb\x03\x5f\xf3\x80\x9c\xbb\x48\x15\xe6\x12\xb6\x82\x06\x6b\x6a\x69\x7b\x4f\x78\xa1\x98\xa6\xe5\x12\xb7\x5c\xfb\x25\xe2\x4d\x41\x7f\xf5\x28\xfa\xfc\x5b\x3a\x09\xaa\xae\x10\x46\x7f\xa4\xa4\xb2\x8a\xad\x87\xd1\xb9\x28\x54\xe0\x42\x6f\x44\x08\xeb\x99\xa1\xa5\xab\x9a\xc1\xe7\xe3\x2a\xae\x5d\x7e\x1d\x56\x63\x54\x83\x7b\x93\xe1\xbd\x94\x12\x1b\x6f\xd9\xb7\x99\x60\xeb\xfe\x3d\xd4\xf7\xf0\x57\x9a\x2e\x34\x0b\x17\x77\xc1\x4e\x61\x5e\xde\xaf\xb3\xe6\xd2\x20\x88\x0c\x69\xdb\x3d\xe0\x56\xde\xdd\x96\x8e\xbb\xd9\x4e\xab\xad\xb4\x87\xb5\x2c\xb5\x89\xc0\xf0\x70\x4c\x70\x06\x79\x12\x98\xcf\xb6\x6e\x5e\x38\xfc\x98\xcd\x4c\xba\x9b\xdf\xad\x2f\x14\x4f\xbe\x70\xef\xbe\xe0\xd4\xcf\x13\x58\x28\xa5\x9d\xd1\xd4\x41\x5e\x4b\x81\x6d\x11\xe7\x53\x9d\xbd\x59\x9d\x5d\x01\xe8\x4b\xca\xed\x66\x9d\xbd\x4b\x21\xdd\x49\x0d\x6d\x49\xed\xb4\x94\xf6\x73\x3e\xba\x8a\xfa\x21\xd4\xd4\x8f\xb8\x9e\xde\xfb\xa3\xaa\xef\xcf\xef\x4e\x90\x7c\x6f\x40\x75\x37\x70\x79\x9b\xb3\xae\x5b\x82\x86\xef\x73\xf0\xd5\x00\x16\xdd\x1f\x7d\xed\x8c\x3c\x76\x41\x1d\xcb\xd1\xf7\xed\x4a\x41\x78\xa8\xf8\xbd\x3b\xec\xbe\x10\x16\x5c\x57\xff\x16\x56\x1b\xb2\x27\x49\xdc\x4a\xea\xb7\xf2\xe8\x27\xcc\xbf\x07\x9e\x67\x7b\xb5\x59\x8b\x2d\x04\x2a\x00\xff\xe4\x46\xef\x0f\x56\xd1\xfd\x86\xd7\xbc\xf5\xa9\x38\x7c\xc3\x1b\x8d\xfb\x35\x04\xd7\x93\xba\x0a\x80\x3a\xff\x83\xa0\xbf\x15\x64\x3b\xee\x27\x29\xb9\x43\x35\x4d\x31\xd4\xde\x73\x87\xe0\xae\xd1\xbf\xa5\xb6\x3b\xe8\xef\x27\xac\xbb\xcd\x37\xbc\xb1\x1c\x3f\x8c\x6b\x67\x6b\xf0\xfa\x32\x76\xba\x7d\x05\x67\x0b\x63\xb0\xdf\x1c\xf1\x18\x40\x7f\xf7\x62\xd8\xaf\x02\xe1\x81\xd8\xce\x56\x55\xc6\x37\xbc\x19\x18\x9e\x6e\x57\x6f\x68\x8c\xb7\xac\xd6\xd8\x05\x9b\xbc\xe7\x68\xd0\xc8\x13\x18\xd9\x02\x8c\x6c\x69\x3c\x5b\xc2\x96\x5d\x4d\x73\x01\xe0\xec\x86\xf6\xdb\xcc\x68\x3b\xac\xdf\x35\x1f\xee\xea\x55\x86\xf9\x12\x54\xdf\xe2\x25\xb7\xe6\xf9\x61\xbb\xce\x4a\xb7\xe8\x60\x0b\xe6\x91\xbb\x4e\xa5\x36\x58\x5b\x1a\x6c\x51\x13\xb8\x37\x20\xab\x27\x00\xff\x30\xfc\xbf\xe1\xbb\xa1\x77\x0f\xf5\xbb\xc3\xf8\x8f\xe7\x35\xd2\x75\xdb\xf2\x77\xfa\x96\xe1\x3f\x13\xee\x77\x2d\x84\xfd\x02\xfb\x0f\xcc\x82\xb6\x02\xfd\x96\xb7\xaf\x69\x52\x79\x9b\xf5\x09\xff\x3f\xe1\xff\xbb\x02\x31\x4f\xf8\xff\x09\xff\x3f\xe1\xff\xef\x8b\xff\xff\x9c\x24\x94\xde\xa6\xf2\x09\xfd\xaf\x45\xff\x46\x56\x1b\x15\x00\x77\xff\xaa\xb9\x21\xae\xa3\x22\xc0\x4d\xf6\x58\xea\x80\x91\xfe\xfa\xbc\xbe\x53\x5b\x7d\xc6\x1d\x82\x4b\x89\x8c\x76\xfe\xde\x8b\x07\xfa\xbe\xf9\xf7\xe2\xfc\x31\x83\xfb\xef\x69\x2d\xeb\x5e\x5a\xff\xf7\xc4\x9c\xdf\x92\x6b\xd1\x3f\x4e\x6e\xc8\xc9\x4b\xb4\xee\x5e\xb4\xef\x41\xe5\x76\xfe\xbf\x27\x8a\x8d\x69\xc8\x18\xd5\x27\x36\xf6\x93\xd4\x5e\xb6\x75\xfd\x17\xca\x86\xd9\x4c\x8f\x8f\xfe\x4f\x9f\x17\xcc\xb7\x2a\x1e\x9e\xde\xa6\x5f\x39\x6a\xbf\xdf\xa6\xb7\xf9\x51\x1b\x44\xdf\xea\xac\x2b\x70\x39\x35\x38\x65\xcb\x17\xab\x37\xf8\x46\x9e\xd5\x6e\x79\x8f\x48\x73\x4b\xe8\xb8\xc5\x8b\xdd\x0d\x25\xd4\x30\x41\xf3\x35\xef\xf6\xbe\x9b\xdb\xf0\x16\x97\xe2\xea\x5d\xc9\xc0\x7a\xde\x8a\x56\x76\x6d\x57\x64\x39\x62\xf5\x57\x0c\x9c\x31\xc5\x76\xf9\x9a\x81\x5d\x7c\xc2\x2f\x78\xbf\x6e\x71\xeb\x7d\x09\xab\x9c\x92\x9d\xee\x33\xde\x3d\xb9\x0d\x2c\x7e\x27\xc1\x1a\x14\x6d\x2b\xa8\x5b\x88\x7a\x13\x69\xad\xfe\x8a\xcf\xda\xeb\x1c\xaf\xae\x31\x2e\x7f\xfb\x8a\x01\x15\x37\xaa\xfc\x1d\x24\xfb\x13\x39\x54\x1d\xe1\x35\xc6\x53\xfd\x48\xff\x54\x4e\x3c\x95\x4a\xe4\x65\x7f\x36\x66\x29\x97\x4a\x77\xdd\xe1\x67\x04\x88\x8e\xf6\x62\x69\x74\xad\x17\xd1\xbf\xe5\xa0\x7f\x50\xe3\xd4\xcf\x1e\xba\x93\x90\xdb\x55\x43\xb4\x76\x47\xb5\x90\x99\xea\x41\x94\x39\x46\x9f\x58\xfe\x84\xc8\x9d\x7e\xa5\xd4\x7e\x17\x2e\x8f\xa6\x28\xb9\xb7\xaf\xc0\x2a\xf5\x3b\xba\x0e\x5a\x72\x5d\x17\x5a\xbe\x13\x0b\xde\x97\xaf\x9c\x7c\xef\x98\x74\x6c\x53\xfa\xd8\x04\x63\x2d\x0f\xff\x3a\x32\xb7\x05\xa1\xc5\x6f\xde\x1f\x0a\xe1\xb5\x24\x91\x16\xa9\x70\xaf\x5f\xc9\xfa\xe5\xc7\x58\x5d\x47\x67\x82\x63\x10\x0e\x5a\x25\x53\xff\x7d\x86\x04\x47\x6c\x9a\xa9\xf6\xae\x23\x96\x49\xf4\xf2\x99\xff\x27\x00\x00\xff\xff\x1a\x5f\x67\x93\xf3\x71\x00\x00"),
          path: "mongo-api.tml",
          root: "mongo-api.tml",
        },
      
        "mongo-solo-readme.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\xcd\x4a\x43\x31\x10\x85\xf7\x79\x8a\x03\xdd\x54\x29\xf7\x01\x04\x37\x56\x85\x2e\x5c\xbb\xed\x34\x99\xe6\x0e\xe6\xce\x94\x64\x6a\xef\xe3\x4b\x5a\x7f\x10\x8a\x66\x33\xe1\x70\xe6\xfb\x60\x5e\x4c\xb3\x3d\x3e\x60\x33\x1d\x0a\x4f\xac\x4e\x2e\xa6\xe1\xfe\xef\x17\xc2\xf5\x3d\x48\x03\x21\x5a\x62\x64\x56\xae\x97\xf0\x40\xf1\x8d\x32\xc3\x47\x72\x1c\xaa\xbd\x4b\xe2\xde\xdb\x51\x93\x08\xf9\x6d\xde\x5b\x85\xa8\x73\xa5\xe8\xa2\x19\x27\xf1\x11\x84\xa3\x26\xae\xa2\x8c\xa9\x8b\xd3\x0e\x89\x9c\x76\xd4\x78\x08\xe1\xd9\x4a\xb1\x53\x2f\x4f\xec\xa3\xa5\x06\xaa\xfc\xc3\xe5\x74\x17\xc2\x62\x81\x57\xf1\x71\xa3\x89\xe7\x10\xb6\xdb\x6d\xb6\xf0\x1d\x2c\xa3\xcf\x88\xa6\xce\xb3\x0f\xeb\xcb\x5c\x21\x5a\x41\xf3\x2a\x9a\x57\x90\x5e\xe3\x86\x61\x18\xa6\x6c\xc3\x79\xeb\x06\x5c\xab\xd5\xcf\xd1\x99\x67\xcd\xd3\xcc\xf1\xcb\xd0\xff\xff\xc3\xf7\x33\xf6\x47\x8d\xcb\x9e\xdd\x76\xfc\xda\x4a\xe1\xd8\xcf\x71\xcd\x11\x3e\x02\x00\x00\xff\xff\x3a\x8b\x2f\x3b\xb4\x01\x00\x00"),
          path: "mongo-solo-readme.tml",
          root: "mongo-solo-readme.tml",
        },
      
        "mongo-solo.tml": { // all .tml assets.
          data: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x4d\x6f\x1b\x37\x10\x3d\xef\xfe\x8a\xa9\x0e\xc1\xae\x21\xac\x73\xea\x41\xad\x0e\x91\xe5\xb4\x39\x24\x2d\xea\x04\x3d\x14\x45\x41\x2d\x67\x25\xa2\x4b\x52\x20\xb9\x96\x0a\xc3\xff\xbd\x18\x7e\xec\x52\x82\x5d\x2b\x45\x91\xd4\x17\x81\x1f\x33\xf3\xe6\xcd\x9b\xe1\xfa\xfa\x1a\xde\x6b\xb5\xd5\xeb\x15\x70\xec\x84\x42\x0b\x0c\x84\x72\x68\x3a\xd6\x22\x1c\x76\xa2\xdd\x01\x1e\xf7\xda\xfa\x13\x89\x6e\xa7\x39\x74\xda\x80\x41\x67\x04\xde\x0b\xb5\x05\x56\x5e\x5f\x83\x24\x37\xcd\x9a\x39\xb6\x61\x16\x81\x29\x1e\xb7\xee\xd0\x5a\xa1\x55\x53\xba\xbf\xf6\x38\x46\x9b\x62\x3c\x94\xc5\x07\x3c\x54\x35\x54\x57\x32\xf3\x30\x07\xbf\x8c\xd6\x73\x40\x63\xb4\xa9\xcb\xc7\x92\x82\xdd\x68\xd5\x89\x2d\xec\x8d\xbe\x17\x1c\x2d\xb4\x7e\x3d\x18\xe6\x84\x56\x1e\x5e\xab\x95\xc2\xd6\x11\x3c\xa7\x81\x01\xdf\x44\x00\xd1\xd4\x3a\x33\xb4\x8e\x82\xff\xa8\xad\x03\xfa\xb3\xce\x08\xb5\x2d\x8b\x37\x83\xdb\xad\x57\xd9\x86\x5f\xe4\x37\x3e\x59\x34\x27\x1b\x3f\x33\x6b\x0f\xda\xf0\x71\xe3\xbd\xe6\xe8\x6f\x50\x0e\xb4\x88\xc0\x7f\x40\x17\x33\x02\xe6\x1c\xca\xbd\xb3\x84\x2f\x92\x89\xe0\x76\x08\x5b\xe1\x59\xb5\xf1\x1e\x65\x13\xb7\x51\xc5\x4c\x9b\xb2\x1b\x54\x9b\x79\xab\xc2\x7e\xcc\x2e\x71\x79\x46\x1e\x65\x2b\x54\xa7\x61\xb1\xf4\xb8\xd6\x82\xf5\xef\x54\xa7\x1f\xca\xa2\x78\xc3\xb9\xb1\x0b\x42\xfc\xdb\xef\x21\x89\x87\x18\x8a\xf8\x79\x9c\x97\x45\xf1\x51\x48\xd4\x83\x5b\x00\x7c\xfb\x1a\xae\xc0\x09\x89\xcd\x1d\xb6\x5a\x71\x3a\x4d\x75\x5b\x24\x88\x81\x46\x3a\x22\xba\x14\x93\xd3\x11\x6d\xd0\x41\xa2\x6d\x3c\x48\x1b\xf3\xb2\x78\x2c\xcb\x82\x2a\x6d\x90\x39\x04\x36\xd2\x11\x24\x29\x99\x50\x8e\x09\x45\xa2\xdc\x6b\xdd\x83\xee\xc0\xea\xf6\x4f\x74\x63\xe5\xb5\xb2\xde\x83\xd3\xa0\x07\x93\x84\xd7\x94\x85\x45\xeb\x19\xc9\x69\xf8\x55\xb8\x1d\x51\x51\xbd\x22\x82\xea\xb2\x10\x9d\xbf\xf2\xcd\x12\x94\xe8\x89\xb8\xc2\xa0\x1b\x8c\xa2\xa5\xb7\x0e\x08\x45\x97\xa0\xfb\x8a\x7f\x0f\xaf\xfd\xdd\x7c\x6f\x19\x35\xa0\xb4\xd3\x4a\xb4\xc1\xce\xa2\x6d\xee\xd0\xd1\x85\x2a\xbb\x3c\x07\x67\x06\xac\xcb\x32\x45\xf3\x58\x95\xe8\xa3\x7c\x3e\xe0\x21\x35\x50\xb8\x40\xf9\x2b\x3c\x80\x50\xd6\x31\xd5\x22\xf1\xc0\x42\xae\x77\x68\xee\xd1\x44\xa5\x4c\x86\xe7\x4a\x49\xfe\x1e\xca\xe2\x9e\x19\x90\x2a\xb4\x6d\xb0\x2e\x0b\xa9\x9a\xd8\x31\xcb\x98\xe9\x04\xee\x95\x54\x11\x57\x66\x92\xcd\x11\xbf\x9b\xd5\x03\x24\x53\x6c\x8b\x24\x67\xe6\x60\x33\x88\x9e\x5b\xb2\x66\x7d\xaf\x0f\x16\x06\xcb\xb6\x31\x81\xd8\x02\xa7\x6d\xed\x34\x6c\x51\xa1\x21\x3d\x50\xce\xde\x3d\xd9\x47\x69\x58\x3f\x74\x78\x9a\x40\x89\x12\x1b\x1b\x3f\xc7\x38\x75\x7f\x48\x6e\xa2\xf7\x8c\xd7\xa4\xba\x13\xcf\x9d\xd1\x32\xef\xd4\x13\x98\x91\xf0\x4a\xc2\x55\x16\xb1\x86\xcb\x67\x1c\xc1\xca\x35\x9a\x75\xb9\x8c\xc5\x78\x51\xa0\xa7\x2a\x9d\xc4\xd4\xac\x57\xa3\x93\x66\xbd\xaa\xe7\xe7\x0a\x3b\x79\x08\x02\x4d\x83\x49\x0f\x41\x9c\xb7\x74\xe9\xe6\x97\x4f\x6b\xd0\x7b\x0c\x49\xfb\x32\x0e\x96\xc8\x08\x45\x67\xd6\x13\x34\x28\x8e\xa6\x17\x0a\xa7\xf9\xbb\x5e\x45\xb7\x0f\x25\x00\xdf\x24\xfd\x95\x40\x8f\x8b\x11\xad\x4d\xbf\xcd\xfb\xf0\xfb\x6c\x6d\x72\xcd\x53\x63\x27\xa1\x57\xf2\xdc\xc5\x1c\x40\xea\x14\xa9\x86\xab\xf5\x8a\x82\x03\x24\x1d\xa7\x35\xfd\xf1\xcd\x02\xa4\x9e\x8f\xeb\xe8\x6a\x01\x32\xec\x3d\x46\x3c\x61\x62\x70\x3c\x02\xdb\xef\x7b\x81\x21\xe1\xc8\x10\x07\xe1\x8f\x6c\x2f\x5a\x24\xe1\x9e\x9c\xb5\xba\xef\x63\x47\x3c\x2d\x1d\xbe\x21\x8c\xf5\x14\xa3\x6a\xdd\x91\xee\x3a\x3c\x3a\xaa\x1e\xfd\xce\xc9\x4f\x7c\x70\xe6\x21\x1e\x5a\x68\x9a\x86\x24\xe5\xad\xea\x20\x28\x08\xc9\x49\x3f\xef\x22\x31\x1f\xf0\xf0\xd1\xb0\x16\xab\xd9\x7a\xd5\x8c\x61\x66\xb5\xbf\xc9\xb1\x43\x03\x92\x6f\x9a\x74\xfd\x56\x0a\x57\xa5\x85\x9f\x93\x67\x76\xf3\x74\x48\x7b\xc1\xb3\x6c\x6e\x15\xaf\xea\xba\x2e\xbd\x53\xd1\x41\x8f\xaa\x8a\x38\x6b\x58\x2e\xfd\xac\x4c\x34\x4f\xe2\x8d\x2c\x07\x24\x63\xa3\xd8\xac\x47\x7c\x22\x7c\xd3\xf0\x4d\xe3\xbb\x2a\xf9\x3f\xe9\x87\xb1\x7e\xcf\xa5\x71\x4b\xdc\x74\xd5\xec\x2d\x13\x3d\x72\x2a\x52\x1b\xde\x9a\xfc\xe1\x8d\x70\xcf\x12\xac\x66\x53\x09\x67\xbe\x0e\xe7\xe7\x9e\xf8\x99\x47\x1b\x02\x79\x22\xce\x92\xa5\xfe\xcc\x93\xf5\xb4\xc7\xe8\xcd\x4d\xaf\x2d\x56\x91\xbc\x4c\x31\x8b\xe5\xc8\x4a\x73\x53\x51\xe8\x70\x85\xd0\xfe\x11\x75\x40\x97\x0c\x53\x5b\x1c\x65\xf1\x50\x8e\xc1\x23\x51\x8b\x65\xe6\xb5\xb9\x55\x76\x30\x18\xb4\xe6\x6d\xea\xef\x9e\xa6\xf3\x33\x29\x45\xef\x77\xa4\x54\x3c\xa1\x96\x97\xc9\x0c\x56\x31\xb7\xcb\x98\x3e\xc1\x7b\x46\x77\x46\xf9\x3f\x66\x13\x74\x7e\x37\xb4\x2d\x62\x68\xe9\x90\x4e\x18\xf9\x63\x41\xfe\x8b\x9c\xea\xa4\x83\xf2\x02\x44\x6f\x85\x12\x76\x87\x1c\x18\xe7\x84\x25\x02\x68\xbc\xed\x33\xd1\x23\x23\x59\x97\x85\x29\x76\x7b\xc4\x76\xfa\x8c\x66\x40\x03\xc8\x4d\x9f\x59\xf1\x65\xa6\xe9\x85\x47\x6c\x07\x7f\xe4\x5f\xe8\x76\xb0\x4e\xcb\xe9\x3e\xdb\xd2\xd7\x98\xf3\x57\x33\x5d\x9d\x4f\x34\x8a\xf7\xf2\x30\xeb\x8e\xde\x31\xa9\x3b\xbc\x90\x37\xa3\xc7\x3a\x3d\x92\x97\x8d\x36\x8a\xf7\x99\x53\x2d\x98\x5c\x34\xd0\x52\x12\xef\xec\xed\x71\x2f\x0c\x72\x4a\xad\xce\xba\x25\x36\x5a\x27\xdd\xd8\x1b\x31\x61\xd8\x31\x4b\xff\x58\x91\xd5\xac\xfe\x37\xd3\x2a\x14\x04\xa7\x47\xf8\xcb\x0c\xa9\xaf\x38\x91\xbf\xf6\x14\xfe\xda\x15\xef\x7c\xe3\xcf\x13\x88\x53\x87\x5f\x80\x9a\xe9\xe5\xe8\x8e\xd5\xd9\x23\xf4\xdc\x7b\xf1\xbf\x17\xf4\x0b\x03\xe1\xa7\x04\x26\xc1\xbb\x80\xea\xa7\x86\xed\xdf\x01\x00\x00\xff\xff\xaa\xf7\x9b\xeb\x69\x11\x00\x00"),
          path: "mongo-solo.tml",
          root: "mongo-solo.tml",
        },
      
    
  }
)

//==============================================================================

// FilesFor returns all files that use the provided extension, returning a
// empty/nil slice if none is found.
func FilesFor(ext string) []string {
  return assets[ext]
}

// MustFindFile calls FindFile to retrieve file reader with path else panics.
func MustFindFile(path string, doGzip bool) (io.Reader, int64) {
  reader, size, err := FindFile(path, doGzip)
  if err != nil {
    panic(err)
  }

  return reader, size
}

// FindDecompressedGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindDecompressedGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, true)
}

// MustFindDecompressedGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindDecompressedGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindDecompressedGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindGzippedFile returns a io.Reader by seeking the giving file path if it exists.
// It returns an uncompressed file.
func FindGzippedFile(path string) (io.Reader, int64, error){
	return FindFile(path, false)
}

// MustFindGzippedFile panics if error occured, uses FindUnGzippedFile underneath.
func MustFindGzippedFile(path string) (io.Reader, int64){
	reader, size, err := FindGzippedFile(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFile returns a io.Reader by seeking the giving file path if it exists.
func FindFile(path string, doGzip bool) (io.Reader, int64, error){
	reader, size, err := FindFileReader(path)
	if err != nil {
		return nil, size, err
	}

	if !doGzip {
		return reader, size, nil
	}

  gzr, err := gzip.NewReader(reader)
	return gzr, size, err
}

// MustFindFileReader returns bytes.Reader for path else panics.
func MustFindFileReader(path string) (*bytes.Reader, int64){
	reader, size, err := FindFileReader(path)
	if err != nil {
		panic(err)
	}
	return reader, size
}

// FindFileReader returns a io.Reader by seeking the giving file path if it exists.
func FindFileReader(path string) (*bytes.Reader, int64, error){
  item, ok := assetFiles[path]
  if !ok {
    return nil,0, fmt.Errorf("File %q not found in file system", path)
  }

  return bytes.NewReader(item.data), int64(len(item.data)), nil
}

// MustReadFile calls ReadFile to retrieve file content with path else panics.
func MustReadFile(path string, doGzip bool) string {
  body, err := ReadFile(path, doGzip)
  if err != nil {
    panic(err)
  }

  return body
}

// ReadFile attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFile(path string, doGzip bool) (string, error){
  body, err := ReadFileByte(path, doGzip)
  return string(body), err
}

// MustReadFileByte calls ReadFile to retrieve file content with path else panics.
func MustReadFileByte(path string, doGzip bool) []byte {
  body, err := ReadFileByte(path, doGzip)
  if err != nil {
    panic(err)
  }

  return body
}

// ReadFileByte attempts to return the underline data associated with the given path
// if it exists else returns an error.
func ReadFileByte(path string, doGzip bool) ([]byte, error){
  reader, _, err := FindFile(path, doGzip)
  if err != nil {
    return nil, err
  }

  if closer, ok := reader.(io.Closer); ok {
    defer closer.Close()
  }

  var bu bytes.Buffer

  _, err = io.Copy(&bu, reader);
  if err != nil && err != io.EOF {
   return nil, fmt.Errorf("File %q failed to be read: %+q", path, err)
  }

  return bu.Bytes(), nil
}
