import { $ } from 'execa'
import fs from 'fs/promises'
import path from 'path'

$.verbose = true

const bin = 'showmd'
const staticTag = 'v1.0.0' // <-- static version/tag
const buildDir = './build'
const goos = ['linux', 'darwin', 'windows']
const goarch = ['amd64', 'arm64']

const name = (GOOS, GOARCH) =>
	`${bin}_${GOOS}_${GOARCH}${GOOS === 'windows' ? '.exe' : ''}`

const filepath = (GOOS, GOARCH) => path.join(buildDir, name(GOOS, GOARCH))

// Ensure build directory exists
await fs.mkdir(buildDir, { recursive: true })

// Download go modules
await $`go mod download`

// Build binaries
await Promise.all(
	goos.flatMap((GOOS) =>
		goarch.map(
			(GOARCH) =>
				$({
					env: { GOOS, GOARCH },
				})`go build -o ${filepath(GOOS, GOARCH)}`
		)
	)
)

console.log(`✅ Built all binaries for tag ${staticTag} in "${buildDir}"`)
