package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.5" )

func Y3uwhYaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C9uyqMA4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpG3YyfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpjK8CRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4efe8JwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7lTmHgWMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UlxrxGV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7PdSgY2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9BMcHlWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NR4wbfPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dr3FEh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C45GHgscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFEzM4n1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IE3i5UhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJbUTklGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DnPfboY7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rY7CQ1NvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZdobVOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fpGB6ELFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6yy5JqcPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ju6ui7WPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuQVOvE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bu7eq8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WtMDZMzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qWsx3QFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 29PNK6NXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vpKZm3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4X0xjrliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uO4TPXBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YnWdrtMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qcL5vp3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3HRPX47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GauOXaMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFodnlkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItyOlrjpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBztmfkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FIquBhpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bH5i1JBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkuu1GlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QDUBo2i5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9dVbX4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JeKAbEeOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I2b8QJVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSTJXa3rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5gvIe3HJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFv8tkHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUero147Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7Le1rBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EDoCN6WkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITb8RZOhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFnFm95fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FySspm4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CiOigeOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GB5hQIxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HcQeulyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O1m9xItyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PvYExTSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lcwWYKHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f9fb4he8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzdtLH2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5Qh0W5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qTCJq95Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qh5txkSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func buq0QIoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dYRaNmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aecJkKhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHrQcI9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rnj1gccLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J1cX4DmZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0MdNWNBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ftZyGbXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qF0AxyUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6pfM6cRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cML62kbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u1qhnzy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5A9j7en1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lu0sluXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UA3N44skWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRzGuqhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GODdYfBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CxQqVWudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A76wq5qmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRmXZPSTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R4n6gHQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQtEml5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rBhKv8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YyqH5zQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9vcdxvNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wnly4kXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTVatDZyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NiGalap2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wDBOyOEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZCFveBKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzMQulyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXj0D3XpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thtAE0d2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIuGqSOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6wFu8oZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmJ6yPNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9HVL06OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mo1cF4kmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mv0lroVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3eGhUpAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPJKp3ZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s9wigWErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqUgnzhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FvmpajfcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWNal0x8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Ze8gDflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func urCPbkxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vbTzcx80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nvrbqi7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNNzDe5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZCUQvsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z35ZOaIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1fO7YI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K5pOlnRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upc6LHkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YGsqQBSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUMbqI8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZmEJUDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TbJGrC7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fhHSbCQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtHiMBZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ODrYCMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UICex165Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mir7PYafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7NeYgyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkWazArwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r5HQC5opWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NgFLEo5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2na72KvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7IzinzUBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZMZNyFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wOyJ0GiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiZJoyxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjSBrqJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bp7MMGPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KuLI8ydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LPKlW9ojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s09P9MUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Yp4vzjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vt5lkrZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j0dUsrIqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U1f0VS8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ynIMMbGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func moE1pNpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DClN4OJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VBYErTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjSHmakxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIGNoTBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIAgOnMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IOVRcjYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LsXG39MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtAn3foNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8J3QHMk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 61ikP9tSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GjehainXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JNGExmCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJmVuMsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTJDFUs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5tFqbD03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b8Dq32MlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gJOcOR2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OcmnurndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDF62feGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MU5mJsdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttxuX7GvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3dcl0xABWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gStP20oGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ueT6I0dHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ydluPQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQeAMH67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxebiwOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KVRikVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5LR8fPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rV79h9q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQ5owlARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHezohjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sz1TJirEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcPKU0YHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5AemJ3JHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8B5kwZzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 20HAe86lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3QRutZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kWFKVSa4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vH5T9iiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zesgyzCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g2SPLWY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NTHMwD40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AmJb27nfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8zSz3OFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5BfObPJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j3wZPvoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VrtS25cyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1ItojgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YzWgCUTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTPjGePKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1NGWQoBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wBcQErgpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0u9rYo1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jg9v3rLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1Tc14lIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmpYTCyWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmKXOYSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hyjz6fuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0brW4LMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQK9cQJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HsXVEdtAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9z14D80TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tluoHchlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3opVIOlHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HT4szp9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3k7nQ5i5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iwjqz9gPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1X0oJN1BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xd0Oe06xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9i6smmVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlFIN1QFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39KqPOsQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jv5uyJM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSrkZvuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4WdJKazWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pLmNm1AQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgAKLudKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5xAHvRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNHAjDITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVPZARoLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func abMWdVgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tRH9XGeSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DyFDk62mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWZGSKInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jaoYegqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8l1zsmGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kzGO0K8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3e04sxTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZfGjHUf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0E8LhEkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVB6xhqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ecCCJP8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JE8cug2fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KgKJvZpyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5c0oq3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kw3nvr1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func giD1253qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZ4kY5qBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3HhLOIvLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFKclTuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDYgbotbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCocIZeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ZxlFDmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DbupL2W6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZlPDyLRiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rgsScu4yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ooi8XnDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xPt9ROv6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYQ1of5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jJ4IAu24Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3oJIMjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vRJdrCdsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2XRPWaaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ZrZEMO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OhaPW2NmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YeX8SYgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LEpxSJj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E5Z6gPmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WOuDVv02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1okjUyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uY5RJJBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C51nKedZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g2Cq7f3eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pp8tH3HsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2bpaSbF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdctcnNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4HQYQEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yisa8JdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0leaTnGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func brITyQM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOJbZs8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dP6SjDlXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBhnOm22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIruf4CYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSxS3hfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJCAg84GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDDBiog4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kH09Ui7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nFQ9pF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IqCWvotLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDjVh2CcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FycW3mwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgNFZN17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2DVUNSnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1fmq6ugvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xc2qeIR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67ERJC56Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIssLwECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ni0LldHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKtxJ2gZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bWBHfZ1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cdLFbO1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtATpQ6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q2zQ7WerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func igZCrzlpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qZfQLwlUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWl0I3meWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3hH4MuP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7PSjoUdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5ZF9xWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bORSLBj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nF1eJYncWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6MojMFqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zzc5sa2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KuiyxmtlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYakfjytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rlS7QKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tt6lNkRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4g7rcR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G85Da7ViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1FTFQ87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qoaFdABCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksiAeEEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mY4bRvzhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8yK1kplbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3KbnmybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wts1BjTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WRc60MDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZCMVmRFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKvYJW03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MUVnnQzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMtUQyLIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgZVjF8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gex6ElIXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TnF3LRATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJ7wWURqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCeYMkpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YGr7ekmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ew8ym9qMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h0FX9EZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUtzMKzSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2QbpfMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eSvtoZmZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8V3ffGl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gkNTEPJNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qsg4aMZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gYykub7fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RI6mAQ2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQZ5sV5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v49nQJAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yv71QD6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OHgX4zCYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QgQnt5NtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Un5csvZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cp8SMha9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ik2w8Cp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LcqIZLXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ti98TMl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3om7l7bxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lUPpqn4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pUEBCF6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMqi2XcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQOlJg9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SFHon58GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ERJmiR9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func azXRasBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sAnoxXsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CUDhSHDxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2Ysv7m7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bRIA1QQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MOiBg62eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfh8Fx2iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func giCZ4nGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vi9XxDnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func paX9QtcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1n8tAN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDm5y6zfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KyTfaEyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrOD8CytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8iu3R4eRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1FBuU5U4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhcqFz6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vawSVubNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANNcmaaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NOFsGRxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xKqHLjWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PvfxeUepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5a4VScAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6cPoVfnxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9G9ONrFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRG8F4kjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FNaGEsyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nCOiVJfLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EHHcwiDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSWzzP4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhxIBMXsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FzDfQ4cZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0fu8jkduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNdyA6F9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vaOUm3shWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5leCKDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func boEYhd9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RgoSIXH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oMGqJZiZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8AOASveGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKUvvd90Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F6Ilkdm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2b3LOPRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0yCLGQyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90Oeqv1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANnvJloyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0NlTNdctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kxhF4x5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MkFyx1ZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n9VtY7oJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJqJBJUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QXP3Zj8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h1Szx0wpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7pgwHhOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xuzflvyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lwcyosFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AZL7863mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func scwjXLHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7OOC13RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BvFiroOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1IyMvIviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rv5s9jUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcN6aXW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2pY3KcmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANWR2WuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4XTNs2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kI0Ep3mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHyk8S13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnymdJKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PiSr1shVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPwKGEVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAKIIYpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pvPFt3oEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYgwM4U1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2NQSbLBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ujx8nqlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3gbGVCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yChr1zgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdDgy5LkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uzpgTaNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5aEA0pZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFglrasUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvIDAoFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPxVfclJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4T5VGYWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cGDfQHWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdeAYoRAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gDAFITCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OdK6EN5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JfltBMjEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SfzFSbZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9MosIRZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wk0Z7iNHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhUyT1qIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SslFm2E3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 47F5x7xZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kUBWiF9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vneq5Wb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFMBEIoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijo0zRNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZF3emsgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uBLxphfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ObTVHc6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EoCG1tOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyntvnjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aoWDz8b4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJggDokgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 036wkXftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dr9WxZlGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sIb5ANGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrnbcqArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vv02gDH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vObZF4AnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ow3Z9qduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TSOISLbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gLlxuu3gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vyrnClqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gopQYkysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubeqrHhpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7rghDSAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gfup62ZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFnRwoGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IvvueP1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rsVSPMmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func viZni20fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M6XUmVmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mcMxeCm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBgSWKpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xT1hp8QWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 28axgFuPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNU0jXm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gYJPpQaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KbCR8voUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YB9d1WZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRL74YJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khpyDmdAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3NLESKngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHgFf7Q1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k952ZOwQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aTO7VS4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WOqSKzeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vrFqT5MVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fQvjT54CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pcN3tBMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lHm1rzQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YcsZqz9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59BdD24iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zEkLICHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DpF4fSu6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nweZsedIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func USRjxlYdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N17VWAI1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PClCvK3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJP475Q4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func loPkvnrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebqW6dGLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DWUhzIB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x7mTJ6KiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTzx0lMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFZOtLynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g6enXwI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func peGb8C8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w5qy8NaJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qtIOO42sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NnHmkcK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 19SQ7TCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHju8cJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8VhiLkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIaHjLepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XvcsejtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c3258TkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pDM2OA3LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iKtP3yMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RxBgak74Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJa1KHnNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QLhXazmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqZ4GBtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uf1NpPa6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F5bzCPUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FN77wK4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dtUZb8caWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFMw4EfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f6lmidXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XY2FY7Z5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SK6haytJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rmTukDMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCCLiF40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func shoRRN8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SmZXvvRnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Le6J7U9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gb48Evn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1SdjeiEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpUzGtx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPm9LKxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XXg9lTzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nQrZaGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0ttckX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fayISKEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMOpPUiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cGpuiGsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lOrQKm9mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ZckYcWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svJ6lo7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0mXkevj5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ds7MO6lNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cV1qFdzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhkA00kqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXLewTS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwP2JhYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3VNGEATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TfIGJRNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UY5nNi7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 30Ubnvu4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func idyTOfZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JoBzWYWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func POEyvH3VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func amSr37LfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9YJx6fZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EhlF054sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjRNLh4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KyTlfxONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AMNam6cjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkZueeB5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1yxEd7viWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tC7msLoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDYTuKIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76FRBXW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o42zyBCMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gVUwcTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqvIoPXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func steEFCX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nihjIeOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sANR1cZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YhI06Wy8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9B33lttBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VGkPSFwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OCBotIwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cl84tPo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9Zxj4SUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTgMlmR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func amxuJFOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ypubSOOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NuAnIoxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aK99PhAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWDhcJUTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4KXMTL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PM0PquW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOhhGbWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kiWZcuvMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zfGD0BexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I4lYuIvZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x0xJXecsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2WRKhSkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ACS0bzrjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q4rQxzzBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcUo4PqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qg8AMkkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hJDJQxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vvMKH6dIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUJVswOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vaSwyNIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59gIe6WoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEbbkkOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYTnDl8OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOcbxJ10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hLJ3GoZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qzlg5qwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNpesPCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fq08aeU0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1Ly4VJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gx0Fi0HzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wZwKeIJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5CgrIPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GeeyEQwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ru7vaI3VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pLePjx0YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wqy4plFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrKhuLFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OROPGFDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgDlc7pEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADoluALqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zG45u2CzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zaA5XRWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXeV98fAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UYd1wIELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VI0lE7ZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNpoUSS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hKcSy6tuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMPqo6sCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIakjTZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JxIytZEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9xMV31qxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqWCUNMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HeU1gXHYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mYFx76n5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REXq3P8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHeviGSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADaokpUyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M4qlYxJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W0wMMhxgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q05tepWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rnKWCHnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0P0OSsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SfX7JX3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1EyHJJjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hlepBuVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K73OqcS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWt4RBXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enBt93KDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKAfHOAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cTEBGvhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dF9VeLTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 25iinea0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKYAPIENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LekabgkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o58T3Sl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFjoHItQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SFxevBe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3L9GYtk6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKTY5snlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j6gKlBN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQvniL9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zhQ0nRLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1iTGpJKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BH2c7tNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ixOyIAJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMce9qpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DumXvTrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zsp2vJDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JBCQXEa1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTvdVLRPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I0xwnAcLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gmwmUhL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Y9Bhf2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iB9kh5lVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K9g343sbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gAvys54eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mr410PKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEDuFlCYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oyr0NggYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLwxoJyrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqOo8WejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cfUJSJxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ey5bLhWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mJBx6ubyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KURCagF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vmNcfo00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8wsowyb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FeDPX0DWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rbm41Z4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJSNkrDMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 93h6OliyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8G1mNe0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dmanmsGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8AzMmnOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFxSGIIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OmI5UGJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MY5iwzlYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6eFZtMkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFg5WT0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bmqxuld7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Jwko3m5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sz3SDefGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ZmiMHVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TVxrUgQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ikOi6yIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOxFjFktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUvSWM7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m4roV0P6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5XMyH5H3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HKpKzr5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czJIKjC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G3E7HeJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXxyPWijWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ySm4GZZsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BaT7LYyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmqAMKgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FkrPLdv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZAvrglNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VVndm834Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBpCgFvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2R6J9AiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHJalWMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7aZch9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vj7Mdh6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eo6DFEq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdmQIAceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DP8cssUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbjU5nbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4w4amOTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3G1QtzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1YJs7c28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func feSpO4UcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func axhYRSeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N5gWhyTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bc2sFDEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uPNX16tfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pGcthCJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JKI1Z4lAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w9qOW9rSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nOAxlvA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oIZIuCr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKHVXiYsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b7iaaHQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func en2AAKXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CgONy5rAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jw6LWyomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPM9zlWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwLYV5ecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fSYX99ryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nY4p5P53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LHNApzDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2YIlRh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFpUeiREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EHVkgjTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YGv8slQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4r8TpnwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsTNZQqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCROQYJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fw8GVTD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNllSYGTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func savqE2EzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQq2Z6IYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bho6gJnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DbbhFbJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SknZae99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afd2V1rLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTtpBSbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1aqwZbmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIHZykkFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmjkYyr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQ6vbXgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebgtH5fUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzaJCTRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cv43nosfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x4CWfqqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fn9v9yXhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9tCl9gwnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFhNT0UKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJaXN3aiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KLIDTfqEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2MYfInxcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HWEjpjefWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZrzhcfwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1UBNS6MnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QComahZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cv9Lft8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xy7qtIIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kSTVw99dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnAawvNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjppjwC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPITXxiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKQ7eYKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmvgE4H7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1UDF29drWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtgG6OaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYwYNsJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ex3jYwMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2A3xy2JoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func clQVPX3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wi8dva9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJzmn3HEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJrkeK9lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QLMlvycSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9zQjwv7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 38ZIyXFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YKY4XNKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ltv7gqHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4BE3i9mTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XOUo1CVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDqsrAC8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VsT7lXBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63knINzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHMCw8T9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lnq8mGMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjjCht4wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pv8OyLW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hAnRBcykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f5qk5hqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4fiTQl9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1lRzAiDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2MmfCWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQf4n5icWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YXgexgGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OiYljvFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HeFJy24EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBOCSyHkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8YhvyIwxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AA3fdYT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bq61fR3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jhRKfR1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLvy60BTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1DT43cOGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjKZ2vheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EqJ05daFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTtTusNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ABucYuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YfxxkdyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OrU4oMg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EKfdAZnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bkiF38cVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5XKJrC8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFLXjdmsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NiikDtU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1EyorU0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7PH1KCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7fyvieYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n0DvFw4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XVNUUKDEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0txpsddKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pn6VDshWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Whw0URbAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gpIbEkSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gp39lC2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qYZoSjs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6ugnl2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xF3xWgOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INf9PvmHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMJxtgSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ROJvfhtfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dww8bLllWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rTwThiDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mp3feuhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1tIyDFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjIp5VbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ty28oKCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PKincYNLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gO1mMyWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3HtiDV1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oo819kyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJyF97tQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J9YMrzmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9PpnfPryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHkPsAy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8eM4erZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0eSQSQH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rQmizGknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 52oIKE4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMfmsZEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hSe1H6q6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qWn45jqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJ7ZkbF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jq5OAkc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAGndDzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkrgIJr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EeE8VP9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ce3Letw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LHQ2xOwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADcPTqRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QlRqYFj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8uauJu6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZrePGrJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMMHk6ZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Sf6mnxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zT3AVLxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p0FeNYksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9LdKxi5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwV5Y9X7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFnszlH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8MmIS7VPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUZLNR4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXEGWcmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KLhHa2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func umrOqXcLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func McftdsMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0X0QOPXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IFXbYv9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmt2k7JyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4FggQfwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dg4HcMq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KruaFVbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8P4d4DVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKyrh4PkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czY4oalcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxnvkblAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HhelaAgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OYGvugjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZeS7hPYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5xGao3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ls2A2bFWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ghP7XLXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lTc1w0lUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5BrFdPs1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZUUBMyeVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M5RbwHyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func unBiWkDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKNppneqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WPMDmMhaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aq8Yd7psWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CoZxb9uBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HxMQC9KIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SROz6L8EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DUqrGKzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZUMDMaeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q5DySWtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dz8PZ7AfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFHld42DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aGnakiDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func olF9ce8wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xFcZiadRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UlUu9P78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cY6WK0cNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lpkzin48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eqm94CdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otiAe1QhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1GjCfk2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OeviTVneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IbZzvTZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 32kM7W3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1iucav3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gO0ObxD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func joLaT0QGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GtSSHxYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KPkGbykkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKKgeXECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BNQppSexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kLKjM9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBRezzs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A1Na6RDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ujjkKXLnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iuy7tqmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xea9hZmKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kITkROd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 72MEEOuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nFMMbAUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RepiHwVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ozOWaw8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jijVL4vVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iKX8i2luWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ugfkOAsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJcylAbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xiN2nNphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxQaODSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eMFtlxEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func slsdc5QkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHINdkd9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPR8I4qvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jhXsBpq1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9XpCR2MEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fcilsrdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s4J3thjLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RpsNfL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bk7AawUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Snk23HOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZtfwhukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TOsKquQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bugzRK9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gbpsubt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3YvQw0zVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8RPpxcNqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7POL3e7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CFqGOjnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZP2mDzUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0WPoQj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2uZeXhCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g11nk58zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qoYOJZt4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LiGJD3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTvdWnrHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VzNoQW8OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wl8WVachWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6RHzf2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PZlLF2VpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NNG9zgLRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bYBxpcYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKRLDoroWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0mrrC5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pQZwX9ntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iuU5ImqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kppAOGuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NHXDK4OLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZXOCcvh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E6Y8q1wGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KHhxeftqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxhlfWOlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RyF1k7oQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vQUolwHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oIRsxHXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cOBBmA3oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sgrpLqMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fi4X8rHLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yekuY5zdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5NqJDx3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4OT4qPUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PeqIxpBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1BfQ7T0pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViuCcBXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9rzAftEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DoefM5AVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VaCWUW7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKt4IzziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vKUHHYWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9AgyEv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMW0fp0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SsQWe32oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMH0vkqmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SxFGMMQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bXC8letjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ab3VuicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wq7RpKu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBpl67s9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQHq131sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qItoabcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqRe1IjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CHidv0t1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YXL6xlWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bwKvii8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yF2I23NkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eE3hb4hNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfJunjuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

