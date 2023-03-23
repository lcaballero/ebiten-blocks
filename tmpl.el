(defun tmpl/join-and-insert (s)
    (let ((dir (file-name-nondirectory
              (directory-file-name
               (file-name-directory buffer-file-name)))))
    (message dir)
    (insert (format (string-join (list s) "\n") dir))))

(defun tmpl/go-new-genart-step ()
  (interactive)
  (tmpl/join-and-insert "\
func (dwg *Polystripe) Step0(dc geom.Context) {
  p := dwg.Store.Params
  rnd := rand.NewRnd(p.Seed)
  r := shapes.Rect{Dim: p.PreviewSize}
  a := r.Shrink(p.Margin)

  dc.Translate(p.PreviewSize.VecY()).FlipY()
  dc.SetHue(colors.VineBlack).DrawRect(a).Stroke()
}
"))

(defun tmpl-go-swap-func ()
  "inserts swap function"
  (interactive)
  (tmpl/join-and-insert "\
  func(i, j int) {
    pts[i], pts[j] = pts[j], pts[i]
  }
"))

(defun tmpl/go-new-simplex ()
  (interactive)
  (tmpl/join-and-insert
   "
  spx0 := geom.NewDefaultSimplexNoise(p.Seed).
    WithScale(shapes.Vec{p.BoxScale, p.BoxScale}).
    WithOffset(shapes.Vec{100, 100})
")
  )

(defun tmpl/go-xy-for-loops ()
  (interactive)
  (tmpl/join-and-insert
   "\
  for x := 0; x < w; x++ {
    for y := 0; y < w; y++ {

    }
  }
")
  )

(defun tmpl-go-new-test-file ()
  "inserts starting code for a new go test file"
  (interactive)
  (let ((dir (file-name-nondirectory
              (directory-file-name
               (file-name-directory buffer-file-name)))))
    (message dir)
    (insert
     (format
      (string-join '("package %s"
                     ""
                     "import ("
                     "  \"testing\""
                     "  \"github.com/stretchr/testify/assert\""
                     ")"
                     ""
                     "func Test_(t *testing.T) {"
                     "  assert.True(t, true)"
                     "}"
                     "")
                   "\n")
      dir))))



(defun tmpl-go-test-cases ()
  "inserts starting code test cases and loop"
  (interactive)
  (tmpl/join-and-insert "\
  cases := []struct {
    name     string
    expected bool
    actual   bool
  }{
    {
      name:     \"starter/basic test\",
      expected: true,
      actual:   true,
    },
  }
  for _, c := range cases {
    t.Run(c.name, func(t *testing.T) {
      assert.True(t, c.expected, c.actual)
    })
  }
")
  )

(defun tmpl-new-bash ()
  "inserts the new bash template"
  (interactive)
  (insert
   (shell-command-to-string
    (format "%s/%s" (getenv "HOME") "dotfiles/bin/templates/tmpl.sh new-bash"))))

(defun tmpl/bash/run-sh ()
  "inserts the script-dir bash code template"
  (interactive)
  (insert (shell-command-to-string (tmpl-this "script-dir"))))

(defun tmpl-go-http-handler-func ()
  "inserts the go-http-handler-func go code template"
  (interactive)
  (insert (shell-command-to-string (tmpl-this "go-http-handler-func"))))


(defhydra my/boilerplates (:color pink :hint nil :exit t)
  "
^Bash^                ^Go^
--------------------------------------------------
_b_: bash file        _h_: go http handler
_d_: bash $DIR        _e_: go test file
^ ^                   _c_: go test cases
^ ^                   _f_: go x by y for loop
^ ^                   _s_: go simplex noise
^ ^                   _w_: go swap at indexes
^ ^                   _t_: go generative step func
"
  ;; bash
  ("b" tmpl/bash/run-sh  nil)
  ("d" tmpl-script-dir nil)

  ;; golang
  ("h" tmpl-go-http-handler-func nil)
  ("e" tmpl-go-new-test-file nil)
  ("c" tmpl-go-test-cases nil)
  ("s" tmpl-go-swap-func nil)
  ("f" tmpl/go-xy-for-loops)
  ("s" tmpl/go-new-simplex)
  ("w" tmpl-go-swap-func)
  ("t" tmpl/go-new-genart-step)

  ;; exit
  ("q" nil "quit" :exit t))

(add-hook 'go-mode-hook
  (lambda ()
   (local-set-key (kbd "C-c C-v") 'my/boilerplates/body)))
