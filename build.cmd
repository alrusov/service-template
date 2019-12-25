@echo off

if exist vendor\github.com (
	echo Please delete the vendor\github.com directory
	echo See https://github.com/golang/go/issues/19000 for details
	exit
)

set CGO_ENABLED=0 
set GOOS=windows

set Y=%DATE:~6,4%
set COPYRIGHT=(C)_2017-%Y%_Alexey_Rusov_(rolic402@mail.ru)

set H1=%TIME:~0,1%
set H2=%TIME:~1,1%
if "%H1%"==" " (set H1=0)
set H=%H1%%H2%

set BUILD_TIME=%Y%-%DATE:~3,2%-%DATE:~0,2%_%H%:%TIME:~3,2%:%TIME:~6,2%

set BUILD=
if exist BUILD_NUMBER (
	for /f %%i in ('type BUILD_NUMBER') do set "BUILD=%%i"
)
if "%BUILD%"=="" (set BUILD=1)

for /f %%i in ('type VERSION') do set "VERSION=%%i"
set VERSION=%VERSION%.%BUILD%

rem Local time, not UTC ((

go build -a --ldflags "-extldflags -static -X github.com/alrusov/misc.appVersion=%VERSION% -X github.com/alrusov/misc.buildTime=%BUILD_TIME% -X github.com/alrusov/misc.copyright=%COPYRIGHT%"

set /a BUILD+=1
echo %BUILD%>BUILD_NUMBER
