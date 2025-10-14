# Stage 1: Build React UI
FROM node:20-alpine AS ui-build
WORKDIR /app/ui
COPY ui/package*.json ./
RUN npm install
COPY ui/ ./
RUN npm run build

# Stage 2: Build Go server
FROM golang:1.25.1-alpine AS go-build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Copy the built React app into server static folder
COPY --from=ui-build /app/ui/dist ./ui/dist
RUN go build -o opsie ./cmd/server

# Stage 3: Run server
FROM alpine:latest
WORKDIR /app
COPY --from=go-build /app/opsie ./
COPY --from=go-build /app/ui/dist ./ui/dist
EXPOSE 8080
CMD ["./opsie"]
