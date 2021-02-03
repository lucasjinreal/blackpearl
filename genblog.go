package main

import (
	"fmt"
	"os/user"
	"time"
	"github.com/abadojack/whatlanggo"
	"os"
	"path/filepath"
	"strings"
	"github.com/daviddengcn/go-colortext"
	"log"
)


type Blog struct {
	Title         string
	Author        string
	CreateTime    string
	Category      string
	Summary       string
	CopyrightInfo string

	SavePath string
}

func IntelligentCategorize(title string, lang whatlanggo.Lang) string {
	if lang == whatlanggo.Eng {
		return "Default Category"
	} else if lang == whatlanggo.Cmn {
		return "默认分类"
	} else {
		return "默认分类"
	}
}

// automatically detect language
func detectLang(title string) whatlanggo.Lang {
	info := whatlanggo.Detect(title)
	lang := info.Lang
	return lang
}

func getCurrentUserName() string {
	currentUserName, err := user.Current()
	if err != nil {
		ct.Foreground(ct.Red, false)
		fmt.Println(err)
	}
	return currentUserName.Username
}
func getNowTimeString() string {
	t := time.Now()
	tFormat := t.Format("2006-01-02 15:05:05")
	return tFormat
}

// this can be change individually
func getCopyrightInfo(lang whatlanggo.Lang) string {
	copyright := ""
	if lang == whatlanggo.Cmn {
		copyright = "> " + "本文由在当地较为英俊的男子金天大神原创，版权所有，欢迎转载，本文首发地址 https://jinfagang.github.io 。" +
			"但请保留这段版权信息，多谢合作，有任何疑问欢迎通过微信联系我交流：`jintianiloveu` \n"
	} else {
		copyright = "> " + "This article was original written by Jin Tian, welcome re-post, first come with https://jinfagang.github.io . " +
			"but please keep this copyright info, thanks, any question could be asked via wechat: `jintianiloveu` \n"
	}
	return copyright
}

// generate summary according to title
func generateSummary(title string, lang whatlanggo.Lang) string {
	pref := ""
	if lang == whatlanggo.Eng {
		pref = "Introduce something about "
	} else {
		pref = "本文介绍 "
	}
	return pref + title
}

// write blog into template
func WriteBlog(blog Blog) {
	if _, err := os.Stat(blog.SavePath); os.IsNotExist(err) {
		// does not exit create path
		os.MkdirAll(blog.SavePath, 777)
		log.Println("directory " + blog.SavePath + " doest not exist, so create it.")
	}
	// do the write thing
	// file name replace space in title, -1 means replace all
	fileName := strings.Replace(blog.Title, " ", "_", -1) + ".md"
	saveFile := filepath.Join(blog.SavePath, fileName)
	ct.Foreground(ct.Green, false)
	fmt.Println("file will be save into : " + saveFile)

	templateContent := "---\n" +
		"title: " + blog.Title + "\n" +
		"date: " + blog.CreateTime + "\n" +
		"category: " + blog.Category + "\n" +
		"---" + "\n" +
		blog.Summary + "\n" +
		"<!-- more -->" + "\n" +
		"# " + blog.Title + "\n" +
		blog.CopyrightInfo + "\n"

	f, err := os.Create(saveFile)
	if err != nil {
		fmt.Println("create file error, " + err.Error())
	} else {
		f.WriteString(templateContent)
	}
}

// write cpp project
func CreateCppProject(projectName string) {
	// this will make a dir in current path named project name
	// and generate main.cpp CMakeLists.txt as well as src/ include/ build/
	os.MkdirAll(projectName, 0700)
	os.MkdirAll(filepath.Join(projectName, "src"), 0700)
	os.MkdirAll(filepath.Join(projectName, "include"), 0700)
	os.MkdirAll(filepath.Join(projectName, "lib"), 0700)
	os.MkdirAll(filepath.Join(projectName, "build"), 0700)

	mCppFile := filepath.Join(projectName, "main.cpp")
	cmakeFile := filepath.Join(projectName, "CMakeLists.txt")

	mCppContent := `/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 * Author: Jin Fagang, All Rights Reserved.
 *
 */


#include <iostream>

/// a simple project
/// \return
int main() {
    std::cout << "Hello, World!" << std::endl;
    return 0;
}`
	cmakeContent := "cmake_minimum_required(VERSION 3.8)\n" +
		"project(" + projectName + ")\nset(CMAKE_CXX_STANDARD 11)\n\n\n" +
			`file(GLOB_RECURSE source_files
				 src/*.cpp
				  lib/*.cpp 
				*.cpp
			src/*.cc
		include/*/*.hpp
	include/*.hpp)

find_package(OpenCV 3.0)
if (NOT OpenCV_FOUND)
    message(FATAL_ERROR "OpenCV > 3.0 not found.")
endif ()

set(BUILD_SHARED_LIBS OFF)` +
		"\n\n\n" +
			"add_executable(main ${source_files})\n" +
				"target_link_libraries(main ${OpenCV_LIBS})"


	f, err := os.Create(mCppFile)
	if err != nil {
		fmt.Println("create file error, " + err.Error())
	} else {
		f.WriteString(mCppContent)
	}

	f2, err := os.Create(cmakeFile)
	if err != nil {
		fmt.Println("create file error, " + err.Error())
	} else {
		f2.WriteString(cmakeContent)
	}
	fmt.Println("cpp project " + projectName + " has been created.")

}
