'use strict';


ngBlogApp.service('blogConfigService',
    function ($log, $http) {
        this.fetchConfig = function (succ) {
            $http({method: 'GET', url: 'api/auth'}).
                success(function (data, status, headers, config) {
                    return succ(data);
                }).
                error(function (data, status, headers, config) {
                    $log.warn(data, status, headers, config);
                });
        };
    });


