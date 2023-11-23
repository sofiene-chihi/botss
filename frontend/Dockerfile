# Stage 1: Build Angular app
FROM node:20-alpine as build

WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files
COPY . .

# Build the Angular app
RUN npm run build -- --output-path=./dist/out

# Stage 2: Serve the built Angular app using Nginx
FROM nginx:alpine

# Remove default nginx website
RUN rm -rf /usr/share/nginx/html/*

# Copy the built Angular app from the previous stage to nginx html directory
COPY --from=build /app/dist/out/browser/ /usr/share/nginx/html
COPY --from=build /app/nginx.config /etc/nginx/nginx.conf

# Copy custom nginx configuration if needed
# COPY nginx.conf /etc/nginx/nginx.conf

# Expose port 80
EXPOSE 80

# Start nginx server
CMD ["nginx", "-g", "daemon off;"]
