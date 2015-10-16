#lang slideshow
(define c (circle 10))
(define r (rectangle 10 20))
(define c_r (hc-append c r))
(define (square n) (filled-rectangle n n))
(hc-append 20 c r c c_r)
(square 10)