// Helpers

function ready(fn) {
	if (document.readyState != "loading") {
		fn();
	} else {
		document.addEventListener("DOMContentLoaded", fn);
	}
}

function outerSizeWithMargin(el) {
	var height = el.offsetHeight;
	var width = el.offsetWidth;
	var style = getComputedStyle(el);

	height += parseInt(style.marginTop) + parseInt(style.marginBottom);
	width += parseInt(style.marginLeft) + parseInt(style.marginRight);
	return { height: height, width: width };
}

// Main

ready(function () {
	console.log("Volatile");
})
