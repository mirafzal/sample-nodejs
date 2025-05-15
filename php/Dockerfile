# Stage 1: Build Laravel app
FROM composer:2 AS build

WORKDIR /app

COPY ./app /app

RUN composer install --no-dev --optimize-autoloader
RUN php artisan config:cache && php artisan route:cache

# Stage 2: Production image with PHP-FPM
FROM php:8.2-fpm-alpine

# Install Laravel dependencies
RUN docker-php-ext-install pdo pdo_mysql

WORKDIR /var/www

# Copy built Laravel app from the builder stage
COPY --from=build /app /var/www

# Permissions
RUN chown -R www-data:www-data /var/www

EXPOSE 9000
CMD ["php-fpm"]
