@echo off

if "%1" == "" (
  echo Use:
  echo %0 COMMAND
  echo where COMMAND one of: start, stop, restart, install, uninstall
  exit 1
)

x86_64\service-template.exe --config ..\..\..\config\service-template-filled.toml --service %1
if %errorlevel% == 109 (
  echo ****
  echo Try to run as administrator
  echo ****
)
