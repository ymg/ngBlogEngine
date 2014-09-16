'use strict';


ngBlogApp.service('authService', function ($rootScope, $log, $http, $q) {

    var userIsAuthenticated = false;

    this.auth = function (up, cb) {
        $http({method: 'POST', url: '/api/login', data: up}).
            success(function (data, status, headers, config) {
                if (status === 200) {
                    userIsAuthenticated = true;
                }
                cb(status)
            }).
            error(function (data, status, headers, config) {
                userIsAuthenticated = false;
                cb(status)
            });
    };
    this.authCheck = function () {
        $http({method: 'PUT', url: '/api/login' }).
            success(function (data, status, headers, config) {
                if (status === 200) {
                    userIsAuthenticated = true;
                } else {
                    userIsAuthenticated = false;
                }
                $rootScope.loginStatus = userIsAuthenticated;
            }).
            error(function (data, status, headers, config) {
                userIsAuthenticated = false;
                $rootScope.loginStatus = userIsAuthenticated;
            });
    };
    this.isLoggedIn = function () {
        return userIsAuthenticated;
    };
    this.logout = function (cb) {
        $http({ method: 'DELETE', url: '/api/logout' }).
            success(function (data, status, headers, config) {
                userIsAuthenticated = false;
                cb(status)
            }).
            error(function (data, status, headers, config) {
                userIsAuthenticated = false;
                cb(status)
            });
    };
});