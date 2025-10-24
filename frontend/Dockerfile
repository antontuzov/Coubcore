# Use the official Node.js image as the base image
FROM node:18-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy the source code into the container
COPY . .

# Build the React application
RUN npm run build

# Use nginx as the production server
FROM nginx:alpine

# Copy the build artifacts from the builder stage
COPY --from=builder /app/build /usr/share/nginx/html

# Copy the nginx configuration
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 80
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]