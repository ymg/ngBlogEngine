'use strict';


ngBlogApp.service('authenticateAdminService',
    function ($log, $http) {

        this.auth = function (up, cb) {
            $http({method: 'POST', url: 'api/auth', data: up}).
                success(function (data, status, headers, config) {
                    cb(status)
                }).
                error(function (data, status, headers, config) {
                    $log.warn(data, status, headers, config);
                });
        };
        this.checkSession = function (f) {
            $http({method: 'POST', url: 'api/authcheck'}).
                success(function (data, status, headers, config) {
                    f(status);
                }).
                error(function (data, status, headers, config) {
                    $log.warn(data, status, headers, config);
                });
        };

    });



