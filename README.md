# Go-Search-Service
An Search Service implemented using Go Lang.

Execution:

- go run hk-search-service.go -h : To know about the input parameters to CLI
- go run hk-search-service.go -query="{your search query}" : To test the service.

Execution (Directly from EXEC - only for mac users):

- change the permission for the executable to work (it was built on macOS Mojave Version:10.14.6)
- ./hk-search-service -h : To know about the input parameters to CLI
- ./hk-search-service -query="{your search query}" : To test the service.

Features:

- It implements LRU-Cache algorithm, and retrieves data from cache for same key searches.
- It stores the local cache in .lrucache.json file in the same directory as the exe/file is executed.

Bug Resolution:

    - If there is any problem with the goquery package. Kindly download the package using the following command:
        - go get -u github.com/PuerkitoBio/goquery
