# Use the node:21.7-alpine image as the base for the build stage
FROM node:21.7-alpine as build

# Set the working directory inside the container
WORKDIR /app/frontend/

# Copy package.json and yarn.lock files to the working directory
COPY ./frontend/package.json ./ 
COPY ./frontend/yarn.lock ./ 

# Install dependencies using yarn
RUN yarn install

# Copy the rest of the frontend source code into the container
COPY ./frontend/ ./

