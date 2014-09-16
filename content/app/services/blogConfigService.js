'use strict';


ngBlogApp.service('blogConfigService', function ($log, $http, $q) {
    this.fetchConfig = function () {

        var deferred = $q.defer();

        $http({method: 'GET', url: '/api/config'}).
            success(function (data, status, headers, config) {
                deferred.resolve(data);
            }).
            error(function (data, status, headers, config) {
                deferred.reject(status);
            });

        return deferred.promise;
    };
});


