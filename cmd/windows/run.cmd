@echo off

call env.cmd
x86_64\service-template.exe --config ..\..\..\config\service-template.toml %*
