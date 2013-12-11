'use strict';


ngBlogApp.controller('AuthCtrl',
    function (authenticateAdminService, $scope) {
        $scope.cred = {}
        $scope.login = function () {
            authenticateAdminService.auth(this.cred);
            $scope.cred = {}
            $scope.loginform.$setPristine();
        }
    });