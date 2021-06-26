resource "google_project_iam_member" "tfer--roles-002F-cloudfunctions-002E-serviceAgentserviceAccount-003A-scoir-002D-318015-0040-appspot-002E-gserviceaccount-002E-com" {
  member  = "serviceAccount:scoir-318015@appspot.gserviceaccount.com"
  project = "scoir-318015"
  role    = "roles/cloudfunctions.serviceAgent"
}

resource "google_project_iam_member" "tfer--roles-002F-editorserviceAccount-003A-scoir-002D-318015-0040-appspot-002E-gserviceaccount-002E-com" {
  member  = "serviceAccount:scoir-318015@appspot.gserviceaccount.com"
  project = "scoir-318015"
  role    = "roles/editor"
}

resource "google_project_iam_member" "tfer--roles-002F-storage-002E-objectAdminserviceAccount-003A-scoir-002D-318015-0040-appspot-002E-gserviceaccount-002E-com" {
  member  = "serviceAccount:scoir-318015@appspot.gserviceaccount.com"
  project = "scoir-318015"
  role    = "roles/storage.objectAdmin"
}

resource "google_project_iam_member" "tfer--roles-002F-storage-002E-objectCreatorserviceAccount-003A-scoir-002D-318015-0040-appspot-002E-gserviceaccount-002E-com" {
  member  = "serviceAccount:scoir-318015@appspot.gserviceaccount.com"
  project = "scoir-318015"
  role    = "roles/storage.objectCreator"
}

resource "google_project_iam_member" "tfer--roles-002F-storage-002E-objectViewerserviceAccount-003A-scoir-002D-318015-0040-appspot-002E-gserviceaccount-002E-com" {
  member  = "serviceAccount:scoir-318015@appspot.gserviceaccount.com"
  project = "scoir-318015"
  role    = "roles/storage.objectViewer"
}
