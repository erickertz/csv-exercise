resource "google_storage_bucket_iam_policy" "tfer--gcf-002D-sources-002D-1026843452605-002D-us-002D-central1" {
  bucket = "b/gcf-sources-1026843452605-us-central1"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "tfer--gcf-002D-sources-002D-1026843452605-002D-us-002D-east1" {
  bucket = "b/gcf-sources-1026843452605-us-east1"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "tfer--gcf-002D-sources-002D-1026843452605-002D-us-002D-west2" {
  bucket = "b/gcf-sources-1026843452605-us-west2"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "tfer--scior-002D-terraform" {
  bucket = "b/scior-terraform"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "tfer--scoir-002D-csv" {
  bucket = "b/scoir-csv"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "tfer--scoir-002D-json" {
  bucket = "b/scoir-json"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketReader"
    },
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyObjectReader"
    }
  ]
}
POLICY
}

resource "google_storage_bucket_iam_policy" "tfer--us-002E-artifacts-002E-scoir-002D-318015-002E-appspot-002E-com" {
  bucket = "b/us.artifacts.scoir-318015.appspot.com"

  policy_data = <<POLICY
{
  "bindings": [
    {
      "members": [
        "projectEditor:scoir-318015",
        "projectOwner:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketOwner"
    },
    {
      "members": [
        "projectViewer:scoir-318015"
      ],
      "role": "roles/storage.legacyBucketReader"
    }
  ]
}
POLICY
}
