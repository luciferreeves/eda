from django.conf import settings
from eda.theme_loader import get_theme_config


def theme_context(request):
    theme = getattr(settings, "ACTIVE_THEME", "default")
    theme_config = get_theme_config()
    return {"ACTIVE_THEME": theme, "THEME_CONFIG": theme_config}
