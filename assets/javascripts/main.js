(function() {
	var readerNoteText = "<p style='position: absolute;'>I often write what I think and what I'm feeling when I think it. This makes for disjointed and tumultuous reading - welcome to a day in my life. I make no apologies for it, but I do hope you'll glean something from it. If anything, I really, REALLY, well and truly, honest to god, hope that someone who reads a post of mine will feel a little better about themselves. Because let's be honest, there's no way in hell anyone actually knows all of the things about all of the things in software enginerding and everyone is incompetent at most of it. If you don't think you are...god help ya. Cheers.</p>";
	$(document).on('ready', function() {
		$('#reader-note').on('click', function(el) {
			if ($(this).children().length) {
				$(this).children().first().remove();
			} else {
				$(this).append(readerNoteText);
			}
		})
	})
})();