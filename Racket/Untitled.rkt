#lang racket

(require "Documents/Dev/testing/learn.rkt")

(require slideshow)

(require slideshow/code)

(define-syntax code+pict
  (syntax-rules ()
    [(code+pict expr)
     (hc-append 10 expr
                (code expr))]))

(require racket/class
         racket/gui/base)

(define f (new frame% [label "My Art"]
                      [width 300]
                      [height 300]
                      [alignment '(center center)]))