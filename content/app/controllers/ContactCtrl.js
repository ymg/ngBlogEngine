'use strict';


ngBlogApp.controller('ContactCtrl',
    function ($scope, postService, $window, authService) {
        $window.document.title = 'Contact me';
    });