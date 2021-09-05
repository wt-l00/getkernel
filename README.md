# getkernel

## what

A tool to search and download kernels for ubuntu


## how to install

``` 
go install github.com/wt-l00/getkernel@latest
```

## list

```
getkernel list
```

## fetch

```
getkernel fetch v5.8.7
```

## example

```
$ getkernel fetch 5.8.7
2020/12/20 11:39:22 downloading: https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.8.7/amd64/linux-headers-5.8.7-050807-generic_5.8.7-050807.202009051031_amd64.deb
2020/12/20 11:39:22 downloading: https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.8.7/amd64/linux-headers-5.8.7-050807_5.8.7-050807.202009051031_all.deb
2020/12/20 11:39:22 downloading: https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.8.7/amd64/linux-image-unsigned-5.8.7-050807-generic_5.8.7-050807.202009051031_amd64.deb
2020/12/20 11:39:22 downloading: https://kernel.ubuntu.com/~kernel-ppa/mainline/v5.8.7/amd64/linux-modules-5.8.7-050807-generic_5.8.7-050807.202009051031_amd64.deb
2020/12/20 11:39:22 Wait for finishes to download
2020/12/20 11:42:12 Finish!!
```

## 参考にした実装

* https://github.com/gfx/get-kernel


