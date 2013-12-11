'use strict';


ngBlogApp.service('authenticateAdminService',
    function ($log, $http) {
        this.auth = function (c) {
            $http({method: 'POST', url: 'api/auth', data: c}).
                success(function (data, status, headers, config) {
                    return status;
                }).
                error(function (data, status, headers, config) {
                    $log.warn(data, status, headers, config);
                });
        };
    });



