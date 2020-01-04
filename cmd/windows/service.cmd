@echo off

if "%1" == "" (
  echo Use:
  echo %0 COMMAND
  echo where COMMAND one of: start, stop, restart, install, uninstall
  exit 1
)

..\..\service-template.exe --config config\service-template-filled.toml --service %1
if %errorlevel% == 13 (
  echo ****
  echo Try to run as administrator
  echo ****
)
