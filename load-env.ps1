param(
    [string]$Path = ".env",
    [switch]$Verbose,
    [switch]$RemoveQuotes
)

$variables = Select-String -Path $Path -Pattern '^\s*[^#\r\n]+=[^\r\n]*$'
foreach($line in $variables) {
    $keyVal = $line.Line -split '=', 2
    $key = $keyVal[0].Trim()
    $val = $RemoveQuotes ? $keyVal[1].Trim('\"').Trim("\'") : $keyVal[1]
    [Environment]::SetEnvironmentVariable($key, $val, 'Process')
    if ($Verbose) {
        Write-Host "Set environment variable: $key=$val"
    }
}
