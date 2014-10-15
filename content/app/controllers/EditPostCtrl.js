'use strict';


ngBlogApp.controller('EditPostCtrl', function ($location, $scope, $q, $log, $window, $routeParams, postService) {
    $window.document.title = 'Edit Post';

    mdarea(jQuery, jQuery.UIkit);

    $scope.postUpdateForm = {};

    if ($routeParams.postId) {
        var markdownEditor = angular.element('.CodeMirror')[0].CodeMirror;
        postService.fetchPost($routeParams.postId)
            .then(function (post) {
                $scope.postUpdateForm.Id = post[0].Id;
                $scope.postUpdateForm.title = post[0].Title;
                markdownEditor.setValue(post[0].Markdown);
            }, function (ignored_status) {
                $log.info('failed fetching post' + ignored_status);
            });
    } else {
        $location.path('/blog');
    }

    $scope.saveChanges = function () {

        if ($scope.postform.$valid) {

            var markdownEditor = angular.element('.CodeMirror')[0].CodeMirror;
            this.postUpdateForm.markdown = markdownEditor.getValue();

            postService.updatePost(this.postUpdateForm).then(function (status) {
                    if (status === 200) {
                        $scope.postUpdateForm = {}
                        markdownEditor.setValue('');
                        $scope.postform.$setPristine();
                        $location.path('/blog/' + $routeParams.postId);
                    } else {
                        $location.path('/login');
                    }
                },
                function (ignored_status) {
                    $.UIkit.notify({
                        message: 'Failed processing your request',
                        timeout: 3000,
                        pos: 'top-center'
                    });
                });
        }
    };
});
