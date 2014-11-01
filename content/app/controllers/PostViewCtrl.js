'use strict';


ngBlogApp.controller('PostViewCtrl', function ($location, $scope, $q, $window, $routeParams, postService, authService) {

    $scope.postView = {};

    $scope.modal = angular.element.UIkit.modal("#confirm-del", {bgclose: false});

    if ($routeParams.postId) {
        postService.fetchPost($routeParams.postId)
            .then(function (data) {

                $scope.postView.Id = data.Id;
                $scope.postView.Title = data.Title;
                $scope.postView.Date = data.Date;
                $scope.postView.Body = data.Body;
                $window.document.title = data.Title;

            }, function (ignored_status) {
                //$console.log(ignored_status);
            });
    } else {
        $location.path('/blog');
    }

    $scope.delete = function (id) {
        postService.deletePost(id).
            then(function (d) {
                $route.reload();
            },
            function (s) {
            });
    };
});
