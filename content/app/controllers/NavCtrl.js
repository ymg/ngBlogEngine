'use strict';


ngBlogApp.controller('NavCtrl', function ($scope, $location, $route) {
    $scope.activePath = null;
    $scope.$on('$routeChangeSuccess', function () {
        $scope.activePath = $location.path();
    });
});