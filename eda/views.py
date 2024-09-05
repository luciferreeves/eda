from eda.theme_loader import render


# Create your views here.
def home(request):
    return render(request, "home.html")
