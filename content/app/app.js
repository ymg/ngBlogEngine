'use strict';


var ngBlogApp = angular.module('ngMicroBlog', ['angular-loading-bar', 'infinite-scroll', 'ngSanitize', 'ngAnimate', 'ngCookies', 'ngRoute'])
    .config(function ($httpProvider, $routeProvider, $locationProvider) {

        $routeProvider.when('/login', {
            templateUrl: '/content/app/directives/login/login.html',
            controller: 'LoginCtrl',
            requireLogin: false
        }).when('/logout', {
            resolve: {
                logout: function ($location, $rootScope, authService) {
                    authService.logout(function (s) {
                        if (s === 200) {
                            $rootScope.loginStatus = authService.isLoggedIn();
                            $location.path('/blog');
                        } else {
                            $.UIkit.notify({
                                message: 'Failed logging out',
                                timeout: 3000,
                                pos: 'top-center'
                            });
                        }

                    });
                }
            },
            requireLogin: true
        }).when('/metaconfig', {
            templateUrl: '/content/app/directives/meta/metacfg.html',
            controller: 'MetaCfgCtrl',
            requireLogin: true
        }).when('/dbconfig', {
            templateUrl: '/content/app/directives/db/dbcfg.html',
            controller: 'DbCfgCtrl',
            requireLogin: true
        }).when('/blog', {
            templateUrl: '/content/app/directives/blog/blog.html',
            controller: 'BlogPostsCtrl',
            requireLogin: false
        }).when('/blog/:postId', {
            templateUrl: '/content/app/directives/post/post.html',
            controller: 'PostViewCtrl',
            requireLogin: false
        }).when('/edit/:postId', {
            templateUrl: '/content/app/directives/postedit/post_editor.html',
            controller: 'EditPostCtrl',
            requireLogin: true
        }).when('/editor', {
            templateUrl: '/content/app/directives/editor/mdedit.html',
            controller: 'BlogPostsEditorCtrl',
            requireLogin: true
        }).when('/contact', {
            templateUrl: '/content/app/directives/contact/contact.html',
            controller: 'ContactCtrl',
            requireLogin: false
        }).otherwise({redirectTo: '/blog'});

        $httpProvider.defaults.xsrfHeaderName = 'X-Csrftoken';
        $locationProvider.html5Mode(true);

        //$locationProvider.hashPrefix('!');

    }).run(function ($rootScope, $cookies, $http, $location, authService) {

        $http.defaults.headers.common['X-Csrftoken'] = atob($cookies._xsrf.split('|')[0]);

        $rootScope.loginStatus = false;
        authService.authCheck();

        $rootScope.highlightCode = function () {
            angular.element(document).ready(function () {
                angular.element('pre code').each(function (i, block) {
                    hljs.highlightBlock(block);
                });
            });
        };

        $rootScope.$on('$routeChangeStart', function (event, next, current) {
            if (next.requireLogin) {
                if (!authService.isLoggedIn()) {
                    $location.path('/login');
                    event.preventDefault();
                }
            }
        });

    }).directive('onFinishRender', function ($timeout) {
        return {
            restrict: 'A',
            link: function (scope, element, attr) {
                if (scope.$last === true) {
                    scope.$evalAsync(attr.onFinishRender);
                }
            }
        }
    });