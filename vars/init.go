package vars

var GlobalFilePathList []string                             // 所有文件的路径
var GlobalFilePath2FileNameExtMap = make(map[string]string) // 文件路径和文件名称（包含后缀）的映射关系
var GlobalFilePath2FileNameMap = make(map[string]string)    // 文件路径和文件名称（不包含后缀）的映射关系
var GlobalFilePath2FileExtMap = make(map[string]string)     // 文件路径和文件扩展名的映射关系

var GlobalHEICFilePathNameMap = make(map[string]string)    // 文件路径和HEIC文件名称的映射关系
var GlobalMOVFilePathPathNameMap = make(map[string]string) // 文件路径和MOV文件名称的映射关系
