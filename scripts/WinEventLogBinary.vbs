' default path is current directory.
Dim objWshShell
Set objWshShell = WScript.CreateObject("WScript.Shell")
path = objWshShell.CurrentDirectory

' arg0 is the destination directory path. (full path)
Dim oParam
set oParam = WScript.Arguments
if oParam.Count > 0 Then
  if InStr(oParam(0), ":") = 0 then
    ' treat arg0 as ralative path
    path = path & "\" & oParam(0)
  Else
    ' treat arg0 as abusolute path
    path = oParam(0)
  end if
end if

' arg1 is the name of event log. (ex. Application, Security, etc..)
if oParam.Count > 1 Then
  logName = oParam(1)
end if

' collecting event logs...
strComputer = "."
Set objWMIService = GetObject("winmgmts:" & "{impersonationLevel=impersonate, (Backup, Security)}!\\" & strComputer & "\root\cimv2")
Set colLogFiles = objWMIService.ExecQuery("SELECT * FROM Win32_NTEventLogFile")
For Each objLogfile in colLogFiles
  if objLogFile.LogFileName = logName Or logName = "" then
    ' If the name of event log was not specified, all of the event logs are saved.
    strBackupLog = objLogFile.BackupEventLog(path & "\" & objLogFile.LogFileName & ".evt")
  end if
Next
