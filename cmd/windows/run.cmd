@echo off

call cmd\windows\env.cmd
%APP% --config config\%APP%.toml %*
