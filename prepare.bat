CD >> tmpvar.txt
set /p GOPATH=<tmpvar.txt
del tmpvar.txt
echo %GO_PATH%
