'use strict';


ngBlogApp.controller('BlogPostsCtrl', function ($scope, $window, $location, $route, $timeout, $log, postService) {

    $window.document.title = 'Blog';
    $scope.busy = false;
    $scope.page = 0;

    $scope.modal = function(id){
        angular.element.UIkit.modal("#confirm-del-"+id, {bgclose: false}).show();
    }

    $scope.loadMore = function () {
        this.busy = true;
        postService.fetchAll(this.page).
            then(function (d) {
                $scope.busy = false;
                $scope.postList = $scope.postList.concat(d);
                $scope.page++;
                $log.info("called");
            },
            function (s) {
                $scope.busy = true;
            });
    };

    $scope.delete = function (id) {
        postService.deletePost(id).
            then(function (d) {
                $route.reload();
            },
            function (s) {
            });
    };

    postService.fetchAll()
        .then(function (currentSet) {
            $scope.postList = currentSet;
            $scope.page++;
        },
        function (ignored_status) {
            //$console.log(ignored_status);
        });
});
