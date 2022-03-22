package media

func (r Resolution) Width() (width int) {
    switch r {
    case Resolution_RESOLUTION_UNSPECIFIED:
        width = 0
    case Resolution_P360:
        width = 480
    case Resolution_P480:
        width = 640
    case Resolution_P720:
        width = 1280
    case Resolution_P1080:
        width = 1920
    case Resolution_K2:
        width = 2048
    case Resolution_K4:
        width = 4096
    case Resolution_K8:
        width = 7680
    case Resolution_ORIGINAL:
        width = 0
    }

    return
}

func (r Resolution) Height() (height int) {
    switch r {
    case Resolution_RESOLUTION_UNSPECIFIED:
        height = 0
    case Resolution_P360:
        height = 360
    case Resolution_P480:
        height = 480
    case Resolution_P720:
        height = 720
    case Resolution_P1080:
        height = 1080
    case Resolution_K2:
        height = 1152
    case Resolution_K4:
        height = 2304
    case Resolution_K8:
        height = 4320
    case Resolution_ORIGINAL:
        height = 0
    }
    return
}
