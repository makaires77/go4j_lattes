function createP1Structure {
    # Pasta raiz do monorepo
    $monorepoPath = Get-Location
    $modulesPath = Join-Path $monorepoPath "pipeline"
    $m1Path = Join-Path $modulesPath "p1_find"

    # Criar a pasta 'modules' se ainda não existir
    if (-Not (Test-Path $modulesPath)) {
        New-Item -ItemType Directory -Path $modulesPath | Out-Null
    }

    # Criar a pasta 'm1_find' se ainda não existir
    if (-Not (Test-Path $m1Path)) {
        New-Item -ItemType Directory -Path $m1Path | Out-Null
    }

    # Lista de subpastas e seus respectivos arquivos
    $subfolders = @{
        "application" = @("csv_handler.go", "search_handler.go", "queue_handler.go")
        "cmd" = @("p1_cmd.go")
        "domain" = @("p1_domain.go")
        "infra" = @("p1_infra.go")
        "interfaces" = @("p1_interface.go", "p1_web.go")
        "tests" = @("p1_tests.go")
    }

    # Criar subpastas e arquivos vazios para a pasta 'm1_find'
    foreach ($subfolder in $subfolders.GetEnumerator()) {
        $subfolderPath = Join-Path $m1Path $subfolder.Key
        if (-Not (Test-Path $subfolderPath)) {
            New-Item -ItemType Directory -Path $subfolderPath | Out-Null
        }

        foreach ($file in $subfolder.Value) {
            $filePath = Join-Path $subfolderPath $file
            if (-Not (Test-Path $filePath)) {
                New-Item -ItemType File -Path $filePath | Out-Null
            }
        }
    }
}

createP1Structure