$ErrorActionPreference = "Stop";
trap { $host.SetShouldExit(1) }

Enable-PSRemoting -Force
