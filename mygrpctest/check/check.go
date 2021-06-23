package check

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	pb "mygrpctest/proto/voice"
	"os"
<<<<<<< HEAD
=======
	"strings"
>>>>>>> d9b048e (second commit)
)

// 定义比较结果返回字段
type Result struct {
	SessionId    string                         `json:"session_id,omitempty"`    // 会话UUID
	StatusCode   int32                          `json:"status_code,omitempty"`   // 评估结果状态码，0：人声；1：静音；2：噪声；3：空包；
	StatusMesg   string                         `json:"status_mesg,omitempty"`   // 评估结果信息
	ModelVersion string                         `json:"model_version,omitempty"` // 模型版本信息
	Sents        []*pb.ProcessResponse_SentInfo `json:"sents,omitempty"`         // 多句子候选
	AudioUrl     string                         `json:"audio_url,omitempty"`     // 语音音频文件下载地址，单词结束时给出
}

<<<<<<< HEAD
// 序列化返回结果
=======
// 序列化outputs的结果
>>>>>>> d9b048e (second commit)
func ProcessRespontoStruct(path string) []Result {
	inputFile, inputError := os.Open(path)
	if inputError != nil {
		return nil
	}
	defer inputFile.Close()
	var s string = "["
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {

			break
		}
		s = s + inputString
	}
<<<<<<< HEAD
	s1 := []byte(s)
	s1[len(s1)-2] = ']'
	s = string(s1)
	// fmt.Printf("all: \n%s", s)
	var resps []Result
	// 直接序列化多个结构体数据
	// 只要json文件结构正确，可以序列化多个
	if err := json.Unmarshal([]byte(s), &resps); err == nil {
=======

	s1 := strings.TrimSpace(s)
	res := []byte(s1)
	res[len(res)-1] = ']'
	var resps []Result
	// 直接序列化多个结构体数据
	// 只要json文件结构正确，可以序列化多个
	if err := json.Unmarshal(res, &resps); err == nil {
>>>>>>> d9b048e (second commit)
		fmt.Println("转换成功")
	} else {
		fmt.Println("转换失败")
	}
	return resps
}

// 序列化返回结果
func ExpectedtoStruct(path string) []Result {
	inputFile, inputError := os.Open(path)
	if inputError != nil {
		return nil
	}
	defer inputFile.Close()
	var s string = "["
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {

			break
		}
		s = s + inputString
	}
<<<<<<< HEAD
	s1 := []byte(s)
	fmt.Println(s1)
	s1[len(s1)-2] = ']'
	s = string(s1)
	//fmt.Printf("all: \n%s", s)
	var resps []Result
	// 直接序列化多个结构体数据
	// 只要json文件结构正确，可以序列化多个
	if err := json.Unmarshal([]byte(s), &resps); err == nil {
		fmt.Println("转换成功")
	} else {
		fmt.Println("转换失败")
=======
	s1 := strings.TrimSpace(s)
	res := []byte(s1)
	res[len(res)-1] = ']'
	// fmt.Printf("all: \n%s", res)
	var resps []Result
	// 直接序列化多个结构体数据
	// 只要json文件结构正确，可以序列化多个
	if err := json.Unmarshal(res, &resps); err == nil {
		fmt.Println("convert successful")
	} else {
		fmt.Println("conver failed")
>>>>>>> d9b048e (second commit)
	}
	return resps
}

<<<<<<< HEAD
=======
//结构体比较
func CheckRes(outputs, expects []Result) bool {
	if len(outputs) != len(expects) {
		fmt.Printf("The expected len is %v,now is %v \n", len(expects), len(outputs))
		return false
	}

	//初始化
	var init, process, close bool

	iniout, expect := outputs[0], expects[0]

	if iniout.StatusMesg == expect.StatusMesg && expect.ModelVersion == expect.ModelVersion {
		init = true
	} else {
		fmt.Printf("the expected is %v , now is %v \n", expect, iniout)
	}
	// 中间处理结果
	iniprocess, exprocess := outputs[1:len(outputs)-1], expects[1:len(expects)-1]
flag:
	for i, v := range iniprocess {
		if v.StatusCode == exprocess[i].StatusCode && v.StatusMesg == exprocess[i].StatusMesg {
			for j, ws := range v.Sents {

				for k, wd := range ws.Words {

					if len(ws.Words) != len(exprocess[i].Sents[j].Words) {
						fmt.Println("the result's lenght is not unanimous")
						process = false
						break flag
					} else {
						if wd.Word == exprocess[i].Sents[j].Words[k].Word {

							for l, ph := range wd.Phones {
								if ph.Phone == exprocess[i].Sents[j].Words[k].Phones[l].Phone && ph.RefPhone == exprocess[i].Sents[j].Words[k].Phones[l].RefPhone {
									process = true
								} else {
									fmt.Printf("the expected is %v , now is %v ,line is %v \n", exprocess[i].Sents[j].Words[k].Phones[l], ph, i)
									process = false
									break flag
								}
							}
							for l, sylls := range wd.Syllables {
								if sylls.Match == exprocess[i].Sents[j].Words[k].Syllables[l].Match && sylls.Syllable == exprocess[i].Sents[j].Words[k].Syllables[l].Syllable {
								} else {
									fmt.Printf("the expected is %v , now is %v,and line is %v \n", exprocess[i].Sents[j].Words[k].Syllables[l], sylls, i)
									process = false
									break flag
								}
							}
							process = true
						} else {
							fmt.Printf("the expected is %v , now is %v \n", wd.Word, exprocess[i].Sents[j].Words[k].Word)
							process = false
							break flag
						}
					}
				}
			}
		} else {
			fmt.Println("StatusCode is not the expected")
		}
	}
	// 关闭通道的结果
	iniclose, expectclose := outputs[len(outputs)-1], expects[len(expects)-1]
	if iniclose.StatusMesg == expectclose.StatusMesg {
		close = true
	} else {
		fmt.Println("initcose is not the expected")
	}
	return init && process && close
}

>>>>>>> d9b048e (second commit)
//获得初始化响应对比字段
func InitRespontoStruct(path string) []pb.InitResponse {
	inputFile, inputError := os.Open(path)
	if inputError != nil {
		return nil
	}
	defer inputFile.Close()
	var s string = "["
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {

			break
		}
		s = s + inputString
	}
	s1 := []byte(s)
	s1[len(s1)-2] = ']'
	s = string(s1)
	//fmt.Printf("all: \n%s", s)
	var resps []pb.InitResponse
	// 直接序列化多个结构体数据
	// 只要json文件结构正确，可以序列化多个
	if err := json.Unmarshal([]byte(s), &resps); err == nil {
		fmt.Println("转换成功")
	} else {
		fmt.Println("转换失败")
	}
	return resps
}

//获得结束响应字段返回结果
func CloseRespontoStruct(path string) []pb.CloseResponse {
	inputFile, inputError := os.Open(path)
	if inputError != nil {
		return nil
	}
	defer inputFile.Close()
	var s string = "["
	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {

			break
		}
		s = s + inputString
	}
	s1 := []byte(s)
	s1[len(s1)-2] = ']'
	s = string(s1)
	//fmt.Printf("all: \n%s", s)
	var resps []pb.CloseResponse
	// 直接序列化多个结构体数据
	// 只要json文件结构正确，可以序列化多个
	if err := json.Unmarshal([]byte(s), &resps); err == nil {
		fmt.Println("转换成功")
	} else {
		fmt.Println("转换失败")
	}
	return resps
}
<<<<<<< HEAD

//结构体比较
func CheckRes(outputs, expects []Result) bool {
	if len(outputs) != len(expects) {
		return false
	}
	//初始化
	var init, process, close bool

	iniout, expect := outputs[0], expects[0]

	if iniout.StatusMesg == expect.StatusMesg && expect.ModelVersion == expect.ModelVersion {
		init = true
	}
	// 中间处理结果
	iniprocess, exprocess := outputs[1:len(outputs)-1], expects[1:len(expects)-1]
flag:
	for i, v := range iniprocess {
		if v.StatusCode == exprocess[i].StatusCode && v.StatusMesg == exprocess[i].StatusMesg {
			for j, ws := range v.Sents {
				for k, wd := range ws.Words {
					if wd.Word == exprocess[i].Sents[j].Words[k].Word {

						for l, ph := range wd.Phones {
							if ph.Phone == exprocess[i].Sents[j].Words[k].Phones[l].Phone && ph.RefPhone == exprocess[i].Sents[j].Words[k].Phones[l].RefPhone {
								process = true
							} else {
								process = false
								break flag
							}
						}
						for l, sylls := range wd.Syllables {
							if sylls.Match == exprocess[i].Sents[j].Words[k].Syllables[l].Match && sylls.Syllable == exprocess[i].Sents[j].Words[k].Syllables[l].Syllable {
							} else {
								process = false
								break flag
							}
						}
						process = true
					} else {
						process = false
						break flag
					}
				}
			}
		}
	}
	// 关闭通道的结果
	iniclose, expectclose := outputs[len(outputs)-1], expects[len(expects)-1]
	if iniclose.StatusMesg == expectclose.StatusMesg {
		close = true
	}
	return init && process && close
}
=======
>>>>>>> d9b048e (second commit)
