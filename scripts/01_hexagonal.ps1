function createHexagonalStructure {
    $monorepoPath = Get-Location
    $projectsPath = Join-Path $monorepoPath "modules"
    $scriptsPath = Join-Path $monorepoPath "scripts"

    # Criar a pasta 'modules' se ainda não existir
    if (-Not (Test-Path $monorepoPath)) {
        New-Item -ItemType Directory -Path $monorepoPath | Out-Null
    }

    # Criar a pasta 'scripts' se ainda não existir
    if (-Not (Test-Path $scriptsPath)) {
        New-Item -ItemType Directory -Path $scriptsPath | Out-Null
    }

    # Criar a estrutura baseada em Ports and Adapters
    $baseStructure = @(
        "application",
        "domain",
        "infra",
        "interfaces",
        "tests",
        "cmd"
    )

    foreach ($projectName in $baseStructure) {
        $projectPath = Join-Path $monorepoPath $projectName
        if (-Not (Test-Path $projectPath)) {
            New-Item -ItemType Directory -Path $projectPath | Out-Null
        }
    }

    # Criar subpastas para os modulos individuais
    $projectNames = @("m1_find", "m2_extract", "m3_analise", "m4_visualize") # Adicione aqui os nomes dos projetos individuais
    foreach ($projectName in $projectNames) {
        $projectPath = Join-Path $projectsPath $projectName

        $applicationPath = Join-Path $projectPath "application"
        $domainPath = Join-Path $projectPath "domain"
        $infraPath = Join-Path $projectPath "infra"
        $interfacesPath = Join-Path $projectPath "interfaces"
        $testsPath = Join-Path $projectPath "tests"
        $cmdPath = Join-Path $projectPath "cmd"

        $folders = @(
            $applicationPath,
            $domainPath,
            $infraPath,
            $interfacesPath,
            $testsPath,
            $cmdPath
        )

        foreach ($folder in $folders) {
            if (-Not (Test-Path $folder)) {
                New-Item -ItemType Directory -Path $folder | Out-Null
            }
        }
    }
}

createHexagonalStructure