package main

var executeJavaScript string = `
    if (document.readyState === "complete" || document.readyState === "interactive") {
      const aktDoms = document.getElementsByClassName('akt');
      let count = 0;

      function computeCount() {
        count = 0;

        for (let i = 0; i < aktDoms.length; i++) {
          const span = aktDoms[i].getElementsByTagName('span')[0];

          count += Number(span.textContent);
        }

        astilectron.sendMessage(count, function() {

        });
      }

      for (let i = 0; i < aktDoms.length; i++) {
        aktDoms[i].addEventListener('DOMSubtreeModified', () => {
          computeCount();
        });
      }

      computeCount();
    } else {
      document.addEventListener('load', function() {
        const aktDoms = document.getElementsByClassName('akt');
        let count = 0;

        function computeCount() {
          count = 0;

          for (let i = 0; i < aktDoms.length; i++) {
            const span = aktDoms[i].getElementsByTagName('span')[0];

            count += Number(span.textContent);
          }

          astilectron.sendMessage(count, function() {

          });
        }

        for (let i = 0; i < aktDoms.length; i++) {
          aktDoms[i].addEventListener('DOMSubtreeModified', () => {
            computeCount();
          });
        }

        computeCount();
      });
    }
`
